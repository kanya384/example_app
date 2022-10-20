package useCase

import (
	"context"
)

type UseCase interface {
	PutFile(ctx context.Context, path string, fileName string, content []byte) (err error)
	DeleteFile(ctx context.Context, path string, fileName string) (err error)
	ListFilesInPath(ctx context.Context, path string) (files []string, err error)
}
