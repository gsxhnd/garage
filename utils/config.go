package utils

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Mode         string `env:"MODE" envDefault:"dev"`
	LogFileName  string `env:"LOG_FILE_NAME" envDefault:"./data/log/garage.log"`
	DatabasePath string `env:"DB_PATH" envDefault:"./data/garage.db"`
	Listen       string `env:"LISTEN" envDefault:":8080"`
	Storage      Storage
}

type Storage struct {
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
	fmt.Println(c)

	return &c, nil
}
