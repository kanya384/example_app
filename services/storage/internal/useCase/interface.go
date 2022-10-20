package useCase

import (
	"context"
)

type UseCase interface {
	PutFile(ctx context.Context, params PutFileParams) (err error)
	PutImageFile(ctx context.Context, params PutImageFileParams) (err error)
	DeleteFile(ctx context.Context, params DeleteFileParams) (err error)
	ListFilesInPath(ctx context.Context, params ListFilesInPathParams) (files []string, err error)
}
