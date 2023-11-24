package types

type AppendObjectOptions struct {
	HashAlgo    HashAlgo
	Hash        []byte
	ContentType string
	Tags        map[string]string
}
