package types

const (
	// VersionLength defines the length for the version field of the bundle.
	VersionLength = 8

	// MetaSizeLength defines the length for the meta size field of the bundle.
	MetaSizeLength = 8

	// TempBundleDir defines the dir under temp dir for storing the temp bundle files.
	TempBundleDir = "/gnfd-bundles"

	// TempBundleFilePrefix defines the prefix for the temp bundle file.
	TempBundleFilePrefix = "bundle-"
)
