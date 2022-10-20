package useCase

import (
	"context"
	"storage/internal/useCase/adapters/storage"
)

type useCase struct {
	storage storage.Storage
	options Options
}

type Options struct {
}

func New(options Options) *useCase {
	var uc = &useCase{}
	uc.SetOptions(options)
	return uc
}

func (uc *useCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func (uc *useCase) PutFile(ctx context.Context, params PutFileParams) (err error) {
	return
}

func (uc *useCase) PutImageFile(ctx context.Context, params PutImageFileParams) (err error) {
	return
}

func (uc *useCase) DeleteFile(ctx context.Context, params DeleteFileParams) (err error) {
	return
}

func (uc *useCase) ListFilesInPath(ctx context.Context, params ListFilesInPathParams) (files []string, err error) {
	return
}
