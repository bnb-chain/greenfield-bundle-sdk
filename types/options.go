package types

// AppendObjectOptions defines the options for appending objects into the bundle.
type AppendObjectOptions struct {
	// HashAlgo indicates the hash algorithm for computing the hash of the object.
	HashAlgo HashAlgo
	// Hash indicates the hash result of the object.
	Hash []byte
	// ContentType indicates the content type of the object.
	ContentType string
	// Tags indicates the tags of the object, each object can be tagged with several attributes.
	Tags map[string]string
}
