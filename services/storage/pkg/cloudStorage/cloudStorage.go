package cloudStorage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
)

type storage struct {
	client *s3.S3
	key    string
}

func NewStorage(key, bucket, endpoint, region string) *storage {
	config := &aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(config))
	client := s3.New(sess, config)

	return &storage{
		client: client,
		key:    key,
	}
}

func (o *storage) PutObjectToStorage(ctx context.Context, bucket, path, fileName string, fileContent io.ReadSeeker) (err error) {
	_, err = o.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path + "/" + fileName),
		Body:   fileContent,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return fmt.Errorf("upload canceled due to timeout, %v", err)
		} else {
			return fmt.Errorf("failed to upload object, %v", err)

		}
	}
	return
}

func (o *storage) ListObjects(ctx context.Context, bucket, prefix string) ([]*string, error) {
	list, err := o.client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	files := make([]*string, len(list.Contents))
	for key, file := range list.Contents {
		files[key] = file.Key
	}
	return files, nil
}

func (o *storage) DeleteObject(ctx context.Context, bucket, path string) error {
	_, err := o.client.DeleteObjectsWithContext(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucket),
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{{Key: aws.String(path)}},
		},
	})
	return err
}
