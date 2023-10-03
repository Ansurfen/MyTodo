package utils

import (
	"encoding/base64"
	"os"
)

func Base64ToFile(src, dst string) error {
	decodeData, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(decodeData)
	return err
}
