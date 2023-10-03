package log

import (
	interfaces "MyTodo/interface"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var _ interfaces.Logger = (*Glog)(nil)

type Glog struct{}

func (y *Glog) Info(msg string) {
	y.logger("INFO", msg)
}

func (y *Glog) Infof(msg string, a ...any) {
	y.logger("INFO", msg, a...)
}

func (y *Glog) Debug(msg string) {
	y.logger("DEBUG", msg)
}

func (y *Glog) Debugf(msg string, a ...any) {
	y.logger("DEBUG", msg, a...)
}

func (y *Glog) Warn(msg string) {
	y.logger("WARN", msg)
}

func (y *Glog) Warnf(msg string, a ...any) {
	y.logger("WARN", msg, a...)
}

func (y *Glog) Error(msg string) {
	y.logger("ERROR", msg)
}

func (y *Glog) Errorf(msg string, a ...any) {
	y.logger("ERROR", msg, a...)
}

func (y *Glog) Fatal(msg string) {
	y.logger("FATAL", msg)
	os.Exit(1)
}

func (y *Glog) Fatalf(msg string, a ...any) {
	y.logger("FATAL", msg, a...)
	os.Exit(1)
}

const defaultTimeFormat = "2006-01-02 15:04:05.000 -0700"

func (y *Glog) logger(level, msg string, a ...any) {
	fr := getTopCaller(3)
	fmt.Printf("%s %s %s:%d %s\n",
		time.Now().Format(defaultTimeFormat),
		level, fr.name, fr.line, fmt.Sprintf(msg, a...))
}

type stackFrame struct {
	name string
	line int
}

func getTopCaller(skip int) stackFrame {
	pc, _, _, _ := runtime.Caller(skip)
	file, line := runtime.FuncForPC(pc).FileLine(pc)
	str := strings.Split(file, "/")
	name := str[len(str)-2] + "/" + str[len(str)-1]
	return stackFrame{
		name: name,
		line: line,
	}
}
