package log

import (
	interfaces "MyTodo/interface"
	"fmt"

	"go.uber.org/zap"
)

var _ interfaces.Logger = (*zlog)(nil)

type zlog struct {
	log *zap.Logger
}

func ZapAdaptor(z *zap.Logger) *zlog {
	return &zlog{log: z}
}

func (z *zlog) Info(msg string) {
	z.log.Info(msg)
}

func (z *zlog) Infof(msg string, v ...any) {
	z.log.Info(fmt.Sprintf(msg, v...))
}

func (z *zlog) Fatal(msg string) {
	z.log.Fatal(msg)
}

func (z *zlog) Fatalf(msg string, v ...any) {
	z.log.Fatal(fmt.Sprintf(msg, v...))
}

func (z *zlog) Debug(msg string) {
	z.log.Debug(msg)
}

func (z *zlog) Debugf(msg string, v ...any) {
	z.log.Debug(fmt.Sprintf(msg, v...))
}

func (z *zlog) Warn(msg string) {
	z.log.Warn(msg)
}

func (z *zlog) Warnf(msg string, v ...any) {
	z.log.Warn(fmt.Sprintf(msg, v...))
}

func (z *zlog) Error(msg string) {
	z.log.Error(msg)
}

func (z *zlog) Errorf(msg string, v ...any) {
	z.log.Error(fmt.Sprintf(msg, v...))
}
