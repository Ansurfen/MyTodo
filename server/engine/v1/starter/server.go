package starter

import (
	"MyTodo/conf"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type MyTodoServer struct {
	*gin.Engine
}

func New(opt Option) *MyTodoServer {
	if len(opt.Mode) > 0 {
		gin.SetMode(opt.Mode)
	}
	s := &MyTodoServer{
		Engine: gin.Default(),
	}
	s.Engine.Use(opt.Middleware...)

	g := s.Group("/")
	for _, f := range opt.Registry {
		f(g)
	}
	return s
}

func (s *MyTodoServer) WaitForShutdownSig(shutdown ...func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
	if len(shutdown) > 0 {
		shutdown[0]()
	}
	os.Exit(1)
}

func (s *MyTodoServer) GET(relativePath string, handlers ...TodoHandler) {
	s.Engine.GET(relativePath, buildHandlers(handlers...)...)
}

func (s *MyTodoServer) POST(relativePath string, handlers ...TodoHandler) {
	s.Engine.POST(relativePath, buildHandlers(handlers...)...)
}

func (s *MyTodoServer) Bootstrap(opt ...conf.ServerOption) error {
	if len(opt) > 0 {
		return s.Engine.Run(fmt.Sprintf("%s:%d", opt[0].Host, opt[0].Port))
	}
	return s.Engine.Run()
}
