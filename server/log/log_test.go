package log_test

import (
	"MyTodo/log"
	"testing"

	"go.uber.org/zap"
)

func TestZLog(t *testing.T) {
	log.Set(log.ZapAdaptor(zap.L()))
	log.Info("Hello World")
}

func TestGlog(t *testing.T) {
	log.Set(&log.Glog{})
	log.Info("Hello World")
	log.Fatalf("%s", "panic!!!")
}
