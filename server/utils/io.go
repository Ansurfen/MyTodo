package utils

import "os"

// Mkdirs recurse to create path
func Mkdirs(path string) error {
	return os.MkdirAll(path, 0777)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
