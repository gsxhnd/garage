package storage

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path"

	"github.com/gsxhnd/garage/utils"
)

type localStorage struct {
	path string
}

func NewLocalStorage(cfg utils.StorageConfig) (Storage, error) {
	if err := utils.MakeDir(cfg.Path); err != nil {
		return nil, err
	}

	if err := utils.MakeDir(path.Join(cfg.Path, "star")); err != nil {
		return nil, err
	}

	if err := utils.MakeDir(path.Join(cfg.Path, "movie")); err != nil {
		return nil, err
	}

	return &localStorage{
		path: cfg.Path,
	}, nil
}

func (s *localStorage) Ping() error {
	return nil
}

func (s *localStorage) GetImage(cover string, id uint, filename string) ([]byte, string, error) {
	var file = path.Join(s.path, cover, "1.jpeg")
	b, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}

	var buf bytes.Buffer
	var tee = io.TeeReader(b, &buf)

	_, f, err := image.Decode(tee)
	if err != nil {
		return nil, "", err
	}

	buff, _ := io.ReadAll(&buf)

	return buff, f, nil
}
