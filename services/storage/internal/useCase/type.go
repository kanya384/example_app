package useCase

type PutFileParams struct {
	path     string
	fileName string
	bucket   string
	content  []byte
}

type PutImageFileParams struct {
	Path      string
	Name      string
	Format    string
	Bucket    string
	MaxWidth  int
	MaxHeight int
	Content   []byte
	Versions  []ImageVersions
}

type ImageVersions struct {
	MaxWidth  int
	MaxHeight int
	Suffix    string
}

type DeleteFileParams struct {
	FilePath string
	Bucket   string
}

type ListFilesInPathParams struct {
	Path   string
	Bucket string
}
