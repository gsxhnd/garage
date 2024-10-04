package storage

import (
	"github.com/google/wire"
	"github.com/gsxhnd/garage/utils"
)

type Storage interface {
	Ping() error
}

func NewStorage(cfg *utils.Config) (Storage, error) {
	if cfg.Storage.Type == "minio" {
		return NewMinioStorage(cfg)
	}
	return &localStorage{}, nil
}

var StorageSet = wire.NewSet(NewStorage)
