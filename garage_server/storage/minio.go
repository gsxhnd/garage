package storage

import (
	"context"
	"errors"

	"github.com/gsxhnd/garage/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioStorage struct {
	Minio *minio.Client
}

func NewMinioStorage(cfg *utils.Config) (*minioStorage, error) {
	minioClient, err := minio.New(cfg.Storage.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	exist, err := minioClient.BucketExists(ctx, cfg.Storage.BucketName)
	if err != nil {
		return nil, err
	}
	if !exist {
		err := minioClient.MakeBucket(ctx, cfg.Storage.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &minioStorage{Minio: minioClient}, nil
}

func (s *minioStorage) Ping() error {
	if s.Minio.IsOffline() {
		return errors.New("minio client offline")
	}
	return nil
}

func (s *minioStorage) GetImage(cover string, id uint, filename string) ([]byte, string, error) {
	return nil, "", nil
}
