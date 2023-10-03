package vfs

import (
	interfaces "MyTodo/interface"
	"MyTodo/utils"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const root = "./assets"

type LocalVFS struct{}

func (fs *LocalVFS) IsLocal() bool {
	return true
}

func (fs *LocalVFS) Mkdir(path string) error {
	return utils.Mkdirs(filepath.Join(root, path))
}

func (fs *LocalVFS) Stat(path string) (fs.FileInfo, error) {
	before, after, ok := strings.Cut(path, ":")
	if !ok {
		return nil, errors.New("invalid path")
	}
	return os.Stat(filepath.Join(root, before, after))
}

func (fs *LocalVFS) Open(path string) (interfaces.VFile, error) {
	path, err := Pathf(path)
	if err != nil {
		return nil, err
	}
	return os.Open(path)
}

func (fs *LocalVFS) IsExist(path string) bool {
	path, err := Pathf(path)
	if err != nil {
		return false
	}
	return utils.IsExist(path)
}

func (fs *LocalVFS) isExist(path string) bool {
	return utils.IsExist(path)
}

func (fs *LocalVFS) Create(path string) (interfaces.VFile, error) {
	return os.Create(path)
}

func (fs *LocalVFS) WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}
