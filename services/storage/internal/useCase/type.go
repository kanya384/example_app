package useCase

type PutFileParams struct {
	path     string
	fileName string
	bucket   string
	content  []byte
}

type PutImageFileParams struct {
	path      string
	name      string
	format    string
	bucket    string
	maxWidth  int
	maxHeight int
	content   []byte
	params    []ImageVersions
}

type ImageVersions struct {
	maxWidth  int
	maxHeight int
	suffix    string
}

type DeleteFileParams struct {
	filePath string
	bucket   string
}

type ListFilesInPathParams struct {
	path   string
	bucket string
}
