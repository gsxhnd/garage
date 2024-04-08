package utils

import (
	"errors"
	"os"
)

func MakeDir(fullPath string) error {
	fs, err := os.Stat(fullPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll(fullPath, os.ModePerm)
		return err
	}

	if !fs.IsDir() {
		return errors.New("file name exist, but not dir")
	}

	if err != nil {
		return err
	}
	return nil
}
