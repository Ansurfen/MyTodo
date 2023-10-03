package starter

import (
	interfaces "MyTodo/interface"
	"MyTodo/utils/vfs"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoContext interface {
	Context() *gin.Context
	Put(k uint8, v any)
	Get(k uint8) any
	ThrowWithResult(err error) (interfaces.Response, error)
	Throw(err error)
	TodoContextParam
	TodoContextReturn
}

type TodoContextReturn interface {
	String(code int, format string, values ...any)
	File(filepath string)
	SaveUploadFile(file *multipart.FileHeader, dst string) error
}

type TodoContextParam interface {
	ParamInt(key string) int
	ParamString(key string) string
	ParamFloat32(key string) float32
	ParamFloat64(key string) float64
}

type Context struct {
	ctx    *gin.Context
	values map[uint8]any
}

func (c *Context) File(filepath string) {
	if vfs.IsLoacl() {
		filepath, err := vfs.Pathf(filepath)
		if err != nil {
			return
		}
		c.ctx.File(filepath)
		return
	}
	file, err := vfs.Open(filepath)
	if err != nil {
		c.Throw(err)
		return
	}
	defer file.Close()
	http.ServeFile(c.ctx.Writer, c.ctx.Request, file.Name())
}

func (c *Context) ParamFloat32(key string) float32 { return 0 }

func (c *Context) ParamFloat64(key string) float64 { return 0 }

func (c *Context) ParamString(key string) string {
	return c.ctx.Param(key)
}

func (c *Context) ParamInt(key string) int {
	if v, err := strconv.Atoi(c.ParamString(key)); err == nil {
		return v
	}
	return 0
}

func (c *Context) Context() *gin.Context {
	return c.ctx
}

func (c *Context) Put(k uint8, v any) {
	if c.values == nil {
		c.values = make(map[uint8]any)
	}
	c.values[k] = v
}

func (c *Context) Get(k uint8) any {
	if c.values == nil {
		c.values = make(map[uint8]any)
	}
	return c.values[k]
}

func (c *Context) Throw(err error) {
	c.ctx.Abort()
	c.ctx.JSON(200, gin.H{
		"data": 500,
		"msg":  err.Error(),
	})
}

func (c *Context) String(code int, format string, a ...any) {
	c.ctx.String(code, format, a...)
}

func (c *Context) ThrowWithResult(err error) (interfaces.Response, error) {
	return &interfaces.ErrorBaseResponse{}, err
}

func (c *Context) SaveUploadFile(file *multipart.FileHeader, dst string) error {
	if vfs.IsLoacl() {
		dst, err := vfs.Pathf(dst)
		if err != nil {
			return err
		}
		return c.ctx.SaveUploadedFile(file, dst)
	}
	fp, err := file.Open()
	if err != nil {
		return err
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return err
	}
	return vfs.WriteFile(dst, data)
}

func UpgradeContext(ctx *gin.Context) TodoContext {
	return &Context{ctx: ctx}
}
