package oss

import (
	"MyTodo/conf"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type TodoMinio struct {
	*minio.Client
}

func New(opt conf.MinioOption) *TodoMinio {
	cli, err := minio.New(
		fmt.Sprintf("%s:%d", opt.Host, opt.Port), &minio.Options{
			Creds:  credentials.NewStaticV4(opt.ID, opt.Secret, ""),
			Secure: opt.Secure,
		})
	if err != nil {
		panic(err)
	}
	return &TodoMinio{cli}
}

func (m *TodoMinio) MakeBucket(name string) error {
	exist, err := m.Client.BucketExists(context.Background(), name)
	if err != nil {
		return err
	}
	if !exist {
		err = m.Client.MakeBucket(context.Background(), name, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *TodoMinio) Put(bucket, filename, path string, opts minio.PutObjectOptions) error {
	_, err := m.Client.FPutObject(context.Background(), bucket, filename, path, opts)
	return err
}
