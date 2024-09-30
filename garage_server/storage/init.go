package storage

import (
	"context"

	"github.com/google/wire"
	"github.com/gsxhnd/garage/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	Minio *minio.Client
}

func NewStorage(cfg *utils.Config) (*Storage, error) {
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

	return &Storage{Minio: minioClient}, nil
}

var StorageSet = wire.NewSet(NewStorage)
