package apis

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/bnb-chain/greenfield-bundle-sdk/types"
)

type Bundle struct {
	version        types.BundleVersion
	metaSize       uint64
	meta           types.BundleMeta
	bundleFileName string
	dataSize       int64

	finalized bool
}

func NewBundle() (*Bundle, error) {
	bundleFile, err := os.CreateTemp("", types.TempBundleFilePrefix)
	defer bundleFile.Close()
	if err != nil {
		return nil, fmt.Errorf("create temp bundle file failed: %v", err)
	}

	return &Bundle{
		version:        types.BundleVersion_V1,
		metaSize:       0,
		meta:           types.BundleMeta{},
		bundleFileName: bundleFile.Name(),
		dataSize:       0,

		finalized: false,
	}, nil
}

func NewBundleFromFile(path string) (*Bundle, error) {
	bundleFile, err := os.Open(path)
	defer bundleFile.Close()
	if err != nil {
		return nil, fmt.Errorf("open bundle failed: %v", err)
	}

	stat, err := bundleFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat bundle failed: %v", err)
	}

	_, err = bundleFile.Seek(types.VersionLength+types.MetaSizeLength, 2)
	if err != nil {
		return nil, fmt.Errorf("seek version and meta size failed: %v", err)
	}

	buf := make([]byte, types.MetaSizeLength+types.VersionLength)
	_, err = bundleFile.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("read version and meta size failed: %v", err)
	}

	version := binary.BigEndian.Uint64(buf[types.MetaSizeLength:])
	if version != uint64(types.BundleVersion_V1) {
		return nil, fmt.Errorf("invalid version")
	}

	metaSize := binary.BigEndian.Uint64(buf[:types.MetaSizeLength])
	if metaSize == 0 {
		return nil, fmt.Errorf("empty bundle")
	}

	buf = make([]byte, metaSize)
	_, err = bundleFile.Seek(int64(metaSize)+types.VersionLength+types.MetaSizeLength, 2)
	if err != nil {
		return nil, fmt.Errorf("seek bundle meta failed: %v", err)
	}
	_, err = bundleFile.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("read bundle meta failed: %v", err)
	}
	bundle := &Bundle{
		version:        types.BundleVersion(version),
		metaSize:       metaSize,
		meta:           types.BundleMeta{},
		bundleFileName: path,
		dataSize:       stat.Size(),
		finalized:      true,
	}
	err = proto.Unmarshal(buf, &bundle.meta)
	if err != nil {
		return nil, fmt.Errorf("unmarshal bundle meta failed: %v", err)
	}

	return bundle, nil
}

func (b *Bundle) AppendObject(name string, size int64, reader io.Reader, options *types.AppendObjectOptions) (*types.ObjectMeta, error) {
	if b.finalized {
		return nil, fmt.Errorf("append not allowed")
	}

	_, err := b.GetObjectMeta(name)
	if err == nil {
		return nil, fmt.Errorf("duplicated name")
	}

	objMeta := &types.ObjectMeta{
		Name:        name,
		Offset:      uint64(b.dataSize),
		Size:        uint64(size),
		HashAlgo:    types.HashAlgo_Unknown,
		Hash:        nil,
		ContentType: "",
		Tags:        nil,
	}

	if options != nil {
		objMeta.HashAlgo = options.HashAlgo
		objMeta.Hash = options.Hash
		objMeta.ContentType = options.ContentType
		objMeta.Tags = options.Tags // map copy here is ok
	}

	bundleFile, err := os.OpenFile(b.bundleFileName, os.O_APPEND|os.O_WRONLY, 0)
	defer bundleFile.Close()
	if err != nil {
		return nil, fmt.Errorf("open bundle failed: %v", err)
	}

	written, err := io.Copy(bundleFile, reader)
	if err != nil {
		return nil, fmt.Errorf("copy to bundle failed: %v", err)
	}
	if written != size {
		bundleFile.Truncate(b.dataSize)
		return nil, fmt.Errorf("written size mismatch, expect: %d, actual: %d", size, written)
	}

	b.dataSize += written
	b.meta.Meta = append(b.meta.Meta, objMeta)
	return objMeta, nil
}

func (b *Bundle) GetObjectMeta(name string) (*types.ObjectMeta, error) {
	for _, objMeta := range b.meta.Meta {
		if objMeta.Name == name {
			return objMeta, nil
		}
	}

	return nil, errors.New("not found")
}

func (b *Bundle) GetObject(name string) (io.Reader, int64, error) {
	objMeta, err := b.GetObjectMeta(name)
	if err != nil {
		return nil, 0, err
	}

	bundleFile, err := os.Open(b.bundleFileName)
	defer bundleFile.Close()
	if err != nil {
		return nil, 0, fmt.Errorf("open bundle failed: %v", err)
	}

	buf := make([]byte, objMeta.Size)
	readBytes, err := bundleFile.ReadAt(buf, int64(objMeta.Offset))
	if err != nil {
		return nil, 0, fmt.Errorf("read object failed: %v", err)
	}
	if uint64(readBytes) != objMeta.Size {
		return nil, 0, fmt.Errorf("object size mismatch, expect: %d, actual: %d", objMeta.Size, readBytes)
	}
	return bytes.NewReader(buf), int64(objMeta.Size), nil
}

func (b *Bundle) FinalizeBundle() (io.ReadCloser, int64, error) {
	if b.finalized {
		return nil, 0, fmt.Errorf("bundle finalized")
	}

	if b.dataSize == 0 {
		return nil, 0, fmt.Errorf("empty bundle")
	}

	bundleFile, err := os.OpenFile(b.bundleFileName, os.O_APPEND|os.O_RDWR, 0)
	if err != nil {
		return nil, 0, fmt.Errorf("open bundle failed: %v", err)
	}

	metaData, err := proto.Marshal(&b.meta)
	if err != nil {
		bundleFile.Close()
		return nil, 0, fmt.Errorf("bundle meta marshal failed: %v", err)
	}

	b.metaSize = uint64(len(metaData))
	buf := make([]byte, types.MetaSizeLength+types.VersionLength)
	binary.BigEndian.PutUint64(buf[:types.MetaSizeLength], b.metaSize)
	binary.BigEndian.PutUint64(buf[types.MetaSizeLength:], uint64(b.version))
	for _, b := range buf {
		metaData = append(metaData, b)
	}

	written, err := bundleFile.Write(metaData)
	if err != nil {
		bundleFile.Close()
		return nil, 0, fmt.Errorf("write bundle meta failed: %v", err)
	}
	if uint64(written) != b.metaSize+types.MetaSizeLength+types.VersionLength {
		bundleFile.Truncate(b.dataSize)
		bundleFile.Close()
		return nil, 0, fmt.Errorf("written size mismatch, expect: %d, actual: %d", b.metaSize+types.MetaSizeLength+types.VersionLength, written)
	}
	b.dataSize += int64(b.metaSize + types.MetaSizeLength + types.VersionLength)
	b.finalized = true
	return bundleFile, b.dataSize, nil
}

func (b *Bundle) GetBundledObject() (io.ReadCloser, int64, error) {
	if !b.finalized {
		return nil, 0, fmt.Errorf("bundle not finalized")
	}

	bundleFile, err := os.Open(b.bundleFileName)
	if err != nil {
		return nil, 0, fmt.Errorf("open bundle failed: %v", err)
	}

	return bundleFile, b.dataSize, nil
}
