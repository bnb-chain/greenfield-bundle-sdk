package bundle

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"

	"github.com/bnb-chain/greenfield-bundle-sdk/types"
)

var (
	objectsNum = 10
	objects    = make([]*singleObject, 0, objectsNum)
)

type singleObject struct {
	name string
	size int64
	data []byte
}

const (
	byteArr       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes   = "abcdefghijklmnopqrstuvwxyz01234569"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = byteArr[rand.Int63()%int64(len(byteArr))]
	}
	return b
}

// randSize generate random int64 between min and max
func randSize() int64 {
	return 64 + rand.Int63n(1024)
}

func randLowerString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// randomObjectName generate random object name.
func randomObjectName() string {
	return randLowerString(10)
}

func setupObjects() error {
	for i := 0; i < objectsNum; i++ {
		name := randomObjectName()
		size := randSize()
		data := randomBytes(int(size))

		objects = append(objects, &singleObject{
			name: name,
			size: size,
			data: data,
		})
	}

	return nil
}

func TestNewBundle(t *testing.T) {
	setupObjects()

	bundle, err := NewBundle()
	require.NoError(t, err)
	defer bundle.Close()

	totoalDataSize := int64(0)
	for _, object := range objects {
		buf := bytes.NewReader(object.data)
		objMeta, err := bundle.AppendObject(object.name, object.size, buf, nil)
		require.NoError(t, err)
		totoalDataSize += int64(objMeta.Size)
	}

	bundledObject, totalSize, err := bundle.FinalizeBundle()
	require.NoError(t, err)
	require.Equal(t, totoalDataSize+int64(bundle.GetBundleMetaSize())+types.MetaSizeLength+types.VersionLength, totalSize)

	bundleFile, err := os.CreateTemp("", types.TempBundleFilePrefix)
	require.NoError(t, err)
	defer bundleFile.Close()

	buf := new(bytes.Buffer)
	readn, err := buf.ReadFrom(bundledObject)
	require.NoError(t, err)
	require.Equal(t, totalSize, readn)

	written, err := bundleFile.Write(buf.Bytes())
	require.NoError(t, err)
	require.Equal(t, totalSize, int64(written))

	bundle2, err := NewBundleFromFile(bundleFile.Name())
	require.NoError(t, err)
	require.Equal(t, totalSize, int64(bundle2.GetBundleSize()))
	require.Equal(t, types.BundleVersion_V1, bundle2.version)
	defer bundle2.Close()

	for _, obj := range objects {
		objMeta := bundle2.GetObjectMeta(obj.name)
		require.NotNil(t, objMeta)
		require.Equal(t, obj.size, int64(objMeta.Size))

		objData, _, err := bundle2.GetObject(obj.name)
		buf = new(bytes.Buffer)
		_, err = buf.ReadFrom(objData)
		require.NoError(t, err)
		require.Equal(t, 0, bytes.Compare(obj.data, buf.Bytes()))
	}
}
