package utils

import (
	"os"
)

func MakeDir(fullPath string) error {
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else if os.IsExist(err) {
			return nil
		} else {
			return err
		}
	}
	return nil
}
