package storage

import (
	"context"
)

type Storage interface {
	PutObjectToStorage(ctx context.Context, path string, fileName string, fileContent []byte) (err error)
	ListObjects(ctx context.Context, prefix string) ([]*string, error)
	DeleteObject(ctx context.Context, path string, fileName string) error
}
