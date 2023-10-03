package vfs

import (
	interfaces "MyTodo/interface"
	"errors"
	"path/filepath"
	"strings"
)

func Set(fs interfaces.TodoFS) {
	vfs = fs
}

func Open(path string) (interfaces.VFile, error) {
	return vfs.Open(path)
}

func Mkdir(path string) error {
	return vfs.Mkdir(path)
}

func IsLoacl() bool {
	return vfs.IsLocal()
}

func IsExist(path string) bool {
	return vfs.IsExist(path)
}

func Create(path string) (interfaces.VFile, error) {
	return vfs.Create(path)
}

func WriteFile(path string, data []byte) error {
	return vfs.WriteFile(path, data)
}

func Copy() {}

func Pathf(path string) (string, error) {
	before, after, ok := strings.Cut(path, ":")
	if !ok {
		return "", errors.New("invalid path")
	}
	return filepath.Join(root, before, after), nil
}

func Objectf(path string) (string, string, error) {
	before, after, ok := strings.Cut(path, ":")
	if !ok {
		return "", "", errors.New("invalid path")
	}
	return before, after, nil
}

var vfs interfaces.TodoFS = &LocalVFS{}
