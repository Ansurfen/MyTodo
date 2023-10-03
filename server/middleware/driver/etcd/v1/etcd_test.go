package etcd_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPathFormat(t *testing.T) {
	fmt.Println(filepath.Join("/todo", "//a"))
	fmt.Println(filepath.Join("/", "/todo", "a"))
	fmt.Println(filepath.Join("//", "todo//", "/a/"))
	fmt.Println(filepath.Join("todo"))
}
