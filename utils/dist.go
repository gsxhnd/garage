package utils

import "os"

func MkdirDistDir(dir_name string) error {
	_, err := os.Stat("./javs")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./javs", os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
