package storage

import (
	"context"
	"io"
)

type Storage interface {
	PutObjectToStorage(ctx context.Context, bucket, path, fileName string, fileContent io.ReadSeeker) (err error)
	ListObjects(ctx context.Context, bucket, prefix string) ([]*string, error)
	DeleteObject(ctx context.Context, bucket, path string) error
}
