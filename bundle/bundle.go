// Package bundle defines the APIs for operating Greenfield bundle.
package bundle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/proto"

	"github.com/bnb-chain/greenfield-bundle-sdk/types"
)

// Bundle indicates the bundle instance for assembling or parsing a Greenfield bundle.
type Bundle struct {
	// version indicates the version of the bundle.
	version types.BundleVersion
	// metaSize indicates the size of the bundle's metadata.
	metaSize uint64
	// meta indicates the metadata for all the objects within the bundle.
	meta types.BundleMeta
	// bundleFile is the file pointer for appending object into the bundle, should not be used in other cases.
	bundleFile *os.File
	// bundleFileName indicates the path of the bundled object.
	bundleFileName string
	// dataSize indicates the size of the bundled object.
	dataSize int64
	// finalized indicates whether the bundle is finalized, once a bundle is finalized, it can't be appended more objects.
	finalized bool
}

// NewBundle creates a new empty bundle with none object bundled.
func NewBundle() (*Bundle, error) {
	bundleFile, err := os.CreateTemp("", types.TempBundleFilePrefix)
	if err != nil {
		return nil, fmt.Errorf("create temp bundle file failed: %v", err)
	}

	return &Bundle{
		version:        types.BundleVersion_V1,
		metaSize:       0,
		meta:           types.BundleMeta{},
		bundleFile:     bundleFile,
		bundleFileName: bundleFile.Name(),
		dataSize:       0,

		finalized: false,
	}, nil
}

// NewBundleFromFile creates a bundle instance for a bundled object.
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

	dataSize := stat.Size()
	_, err = bundleFile.Seek(dataSize-(types.VersionLength+types.MetaSizeLength), 0)
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
	_, err = bundleFile.Seek(dataSize-(int64(metaSize)+types.VersionLength+types.MetaSizeLength), 0)
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
		bundleFile:     nil,
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

// AppendObject is used for appending a new object into the non-finalized bundle.
func (b *Bundle) AppendObject(name string, size int64, reader io.Reader, options *types.AppendObjectOptions) (*types.ObjectMeta, error) {
	if b.finalized {
		return nil, fmt.Errorf("append not allowed")
	}

	objMeta := b.GetObjectMeta(name)
	if objMeta != nil {
		return nil, fmt.Errorf("duplicated name")
	}

	objMeta = &types.ObjectMeta{
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

	written, err := io.Copy(b.bundleFile, reader)
	if err != nil {
		return nil, fmt.Errorf("copy to bundle failed: %v", err)
	}
	if written != size {
		b.bundleFile.Truncate(b.dataSize)
		return nil, fmt.Errorf("written size mismatch, expect: %d, actual: %d", size, written)
	}

	b.dataSize += written
	b.meta.Meta = append(b.meta.Meta, objMeta)
	return objMeta, nil
}

// GetObjectMeta returns the metadata of the specified object.
func (b *Bundle) GetObjectMeta(name string) *types.ObjectMeta {
	for _, objMeta := range b.meta.Meta {
		if objMeta.Name == name {
			return objMeta
		}
	}

	return nil
}

// GetObject returns the object content from the bundled object.
func (b *Bundle) GetObject(name string) (io.ReadCloser, int64, error) {
	objMeta := b.GetObjectMeta(name)
	if objMeta == nil {
		return nil, 0, fmt.Errorf("object not found")
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
	return ioutil.NopCloser(bytes.NewReader(buf)), int64(objMeta.Size), nil
}

// FinalizeBundle is used to finalize a bundle, once the bundle is finalized, it can't be appended more objects.
func (b *Bundle) FinalizeBundle() (io.ReadSeekCloser, int64, error) {
	if b.finalized {
		return nil, 0, fmt.Errorf("bundle finalized")
	}

	if b.dataSize == 0 {
		return nil, 0, fmt.Errorf("empty bundle")
	}

	metaData, err := proto.Marshal(&b.meta)
	if err != nil {
		return nil, 0, fmt.Errorf("bundle meta marshal failed: %v", err)
	}

	b.metaSize = uint64(len(metaData))
	buf := make([]byte, types.MetaSizeLength+types.VersionLength)
	binary.BigEndian.PutUint64(buf[:types.MetaSizeLength], b.metaSize)
	binary.BigEndian.PutUint64(buf[types.MetaSizeLength:], uint64(b.version))
	metaData = append(metaData, buf...)

	written, err := b.bundleFile.Write(metaData)
	if err != nil {
		return nil, 0, fmt.Errorf("write bundle meta failed: %v", err)
	}
	if uint64(written) != b.metaSize+types.MetaSizeLength+types.VersionLength {
		b.bundleFile.Truncate(b.dataSize)
		return nil, 0, fmt.Errorf("written size mismatch, expect: %d, actual: %d", b.metaSize+types.MetaSizeLength+types.VersionLength, written)
	}

	_, err = b.bundleFile.Seek(0, 0)
	if err != nil {
		b.bundleFile.Truncate(b.dataSize)
		return nil, 0, fmt.Errorf("seek to bundle start failed: %v", err)
	}

	b.dataSize += int64(b.metaSize + types.MetaSizeLength + types.VersionLength)
	b.finalized = true
	return b.bundleFile, b.dataSize, nil
}

// GetBundledObject returns the bundled object once the bundle is finalized.
func (b *Bundle) GetBundledObject() (io.ReadSeekCloser, int64, error) {
	if !b.finalized {
		return nil, 0, fmt.Errorf("bundle not finalized")
	}

	bundleFile, err := os.Open(b.bundleFileName)
	if err != nil {
		return nil, 0, fmt.Errorf("open bundle failed: %v", err)
	}

	return bundleFile, b.dataSize, nil
}

// GetBundleMetaSize returns the metadata size of the bundle.
func (b *Bundle) GetBundleMetaSize() uint64 {
	return b.metaSize
}

// GetBundleObjectsMeta returns the objects' metadata within the bundle.
func (b *Bundle) GetBundleObjectsMeta() []*types.ObjectMeta {
	return b.meta.GetMeta()
}

// GetBundleSize returns the size of the bundled object.
func (b *Bundle) GetBundleSize() uint64 {
	return uint64(b.dataSize)
}

// GetBundleVersion returns the version of the bundle.
func (b *Bundle) GetBundleVersion() types.BundleVersion {
	return b.version
}
