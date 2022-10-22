package useCase

import (
	"bytes"
	"context"
	"storage/internal/useCase/adapters/imageResizer"
	"storage/internal/useCase/adapters/storage"
	"storage/pkg/helpers"
	"sync"
)

type useCase struct {
	storage      storage.Storage
	imageResizer imageResizer.ImageResizer
	options      Options
}

type Options struct {
}

func New(storage storage.Storage, imageResizer imageResizer.ImageResizer, options Options) *useCase {
	var uc = &useCase{
		storage:      storage,
		imageResizer: imageResizer,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *useCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func (uc *useCase) PutFile(ctx context.Context, params PutFileParams) (err error) {
	return uc.storage.PutObjectToStorage(ctx, params.bucket, params.path, params.fileName, bytes.NewReader(params.content))
}

func (uc *useCase) PutImageFile(ctx context.Context, params PutImageFileParams) (err error) {
	// TODO - продумать rollback (удаление созданных файлов, если на каком-то этапе ошибка)
	mainImg, err := uc.imageResizer.ResizeImage(params.Content, params.MaxWidth, params.MaxHeight)
	if err != nil {
		return
	}

	imageBuf, err := helpers.ConvertImageToBytes(mainImg)
	if err != nil {
		return
	}

	imageReader := bytes.NewReader(imageBuf)

	err = uc.storage.PutObjectToStorage(ctx, params.Bucket, params.Path, params.Name+"."+params.Format, imageReader)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}

	for _, version := range params.Versions {
		wg.Add(1)
		go func(imageBytes []byte, name string, version ImageVersions) {
			defer wg.Done()

			img, err := uc.imageResizer.ResizeImage(params.Content, version.MaxWidth, version.MaxHeight)
			if err != nil {
				return
			}

			imageBuf, err := helpers.ConvertImageToBytes(img)
			if err != nil {
				return
			}

			imageReader := bytes.NewReader(imageBuf)

			err = uc.storage.PutObjectToStorage(ctx, params.Bucket, params.Path, params.Name+version.Suffix+"."+params.Format, imageReader)
			if err != nil {
				return
			}
		}(params.Content, params.Name, version)
	}
	wg.Wait()
	return
}

func (uc *useCase) DeleteFile(ctx context.Context, params DeleteFileParams) (err error) {
	return uc.storage.DeleteObject(ctx, params.Bucket, params.FilePath)
}

func (uc *useCase) ListFilesInPath(ctx context.Context, params ListFilesInPathParams) (files []*string, err error) {
	return uc.storage.ListObjects(ctx, params.Bucket, params.Path)
}
