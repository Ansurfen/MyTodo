package interfaces

type TodoFS interface {
	Mkdir(path string) error
	Open(path string) (VFile, error)
	IsLocal() bool
	IsExist(path string) bool
	Create(path string) (VFile, error)
	WriteFile(path string, data []byte) error
}

type VFile interface {
	Name() string
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
}
