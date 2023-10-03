package vfs

import (
	interfaces "MyTodo/interface"
	"MyTodo/middleware/driver/oss/v1"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

func OSSAdapater(o *oss.TodoMinio) interfaces.TodoFS {
	return &OSSVFS{cli: o}
}

type OSSVFS struct {
	LocalVFS
	cli *oss.TodoMinio
}

func NewOSSVFS(o *oss.TodoMinio) *OSSVFS {
	return &OSSVFS{cli: o}
}

func (fs *OSSVFS) IsLocal() bool {
	return false
}

func (fs *OSSVFS) Copy() {}

var _ os.FileInfo = (*OSSFileInfo)(nil)

type OSSFileInfo struct {
	bucket string
	*minio.ObjectInfo
}

func (info *OSSFileInfo) Name() string {
	return fmt.Sprintf("%s:%s", info.bucket, info.Key)
}

func (info *OSSFileInfo) Size() int64 {
	return info.ObjectInfo.Size
}

func (info *OSSFileInfo) Mode() os.FileMode {
	return 0
}

func (info *OSSFileInfo) ModTime() time.Time {
	return info.LastModified
}

func (info *OSSFileInfo) IsDir() bool {
	return false
}

func (info *OSSFileInfo) Sys() any {
	return info
}

func (info *OSSFileInfo) String() string {
	data, _ := json.Marshal(info.ObjectInfo)
	return string(data)
}

func (fs *OSSVFS) Mkdir(path string) error {
	return fs.cli.MakeBucket(path)
}

func (fs *OSSVFS) Stat(path string) (os.FileInfo, error) {
	bucket, obejct, err := Objectf(path)
	if err != nil {
		return nil, err
	}
	info, err := fs.cli.GetObjectACL(context.Background(), bucket, obejct)
	return &OSSFileInfo{ObjectInfo: info}, err
}

func (fs *OSSVFS) Open(path string) (interfaces.VFile, error) {
	bucket, obejct, err := Objectf(path)
	if err != nil {
		return nil, err
	}
	cachePath, err := Pathf(fmt.Sprintf("oss:%s/%s", bucket, obejct))
	if err != nil {
		return nil, err
	}
	if fs.LocalVFS.isExist(cachePath) {
		return os.Open(cachePath)
	}
	err = fs.cli.Client.FGetObject(context.TODO(), bucket, obejct, cachePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &OSSFile{path: cachePath}, nil
}

func (fs *OSSVFS) Create(path string) (interfaces.VFile, error) {
	bucket, object, err := Objectf(path)
	if err != nil {
		return nil, err
	}
	_, err = fs.cli.PutObject(context.Background(), bucket, object, strings.NewReader(""), 0, minio.PutObjectOptions{})
	return &OSSFile{bucket: bucket, object: object, cli: fs.cli}, err
}

func (fs *OSSVFS) IsExist(path string) bool {
	bucket, obejct, err := Objectf(path)
	if err != nil {
		return false
	}
	_, err = fs.cli.Client.GetObjectACL(context.Background(), bucket, obejct)
	return err == nil
}

func (fs *OSSVFS) WriteFile(path string, data []byte) error {
	bucket, object, err := Objectf(path)
	if err != nil {
		return err
	}
	_, err = fs.cli.PutObject(context.Background(), bucket, object, bytes.NewBuffer(data), int64(len(data)), minio.PutObjectOptions{})
	return err
}

type OSSFile struct {
	bucket string
	object string
	path   string
	file   *os.File
	cli    *oss.TodoMinio
}

func (f *OSSFile) Name() string {
	return f.path
}

func (f *OSSFile) Read(b []byte) (n int, err error) {
	if f.file == nil {
		f.file, err = os.Open(f.path)
		if err != nil {
			return
		}
	}
	return f.file.Read(b)
}

func (f *OSSFile) Write(b []byte) (n int, err error) {
	_, err = f.cli.PutObject(context.Background(), f.bucket, f.object, bytes.NewReader(b), int64(len(b)), minio.PutObjectOptions{})
	return
}

func (f *OSSFile) Close() error {
	return nil
}

// upload
func (fs *OSSVFS) Write() {}

func (fs *OSSVFS) ReadDir() {
	fs.cli.Client.ListBuckets(context.Background())
}

// user:/profile/
func (fs *OSSVFS) Walk(path string) error {
	before, after, ok := strings.Cut(path, ":")
	if !ok {
		return errors.New("invalid path")
	}
	objs := fs.cli.Client.ListObjects(context.Background(), before, minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    after,
	})
	for obj := range objs {
		fmt.Println(obj)
	}
	return nil
}
