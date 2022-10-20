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

func (uc *useCase) PutFile(ctx context.Context, path string, fileName string, content []byte) (err error) {
	return
}

func (uc *useCase) DeleteFile(ctx context.Context, path string, fileName string) (err error) {
	return
}

func (uc *useCase) ListFilesInPath(ctx context.Context, path string) (files []string, err error) {
	return
}
