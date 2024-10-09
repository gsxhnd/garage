package utils

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Mode         string `env:"MODE" envDefault:"dev"`
	DatabasePath string `env:"DB_PATH" envDefault:"./data/garage.db"`
	Listen       string `env:"LISTEN" envDefault:":8080"`
	Log          LogConfig
	Storage      StorageConfig
}

type LogConfig struct {
	Level      string `env:"LOG_LEVEL" envDefault:"debug"`
	FileName   string `env:"LOG_FILE_NAME" envDefault:"./data/log/garage.log"`
	MaxBackups int    `env:"LOG_MAX_BACKUPS" envDefault:"10"`
	MaxAge     int    `env:"LOG_MAX_AGE" envDefault:"7"`
}

type StorageConfig struct {
	Type       string `env:"STORAGE_TYPE" envDefault:"local"`
	Path       string `env:"STORAGE_PATH" envDefault:"./data/cover"`
	Endpoint   string `env:"STORAGE_ENDPOINT" envDefault:"localhost:9000"`
	BucketName string `env:"STORAGE_BUCKET_NAME" envDefault:"jav-cover"`
	AccessKey  string `env:"STORAGE_ACCESS_KEY"`
	SecretKey  string `env:"STORAGE_SECRET_KEY"`
}

func NewConfig() (*Config, error) {
	var c Config

	opts := env.Options{
		Prefix: "GARAGE_",
	}

	if err := env.ParseWithOptions(&c, opts); err != nil {
		return nil, err
	}

	return &c, nil
}
