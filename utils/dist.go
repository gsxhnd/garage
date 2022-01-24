package utils

import "os"

func MkdirDistDir(dir_name string) error {
	_, err := os.Stat(dir_name)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir_name, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
