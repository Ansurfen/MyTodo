package log

import interfaces "MyTodo/interface"

var logger interfaces.Logger = &Glog{}

func Set(l interfaces.Logger) {
	logger = l
}

func Info(msg string) {
	logger.Info(msg)
}

func Infof(msg string, v ...any) {
	logger.Infof(msg, v...)
}

func Fatal(msg error) {
	logger.Fatal(msg.Error())
}

func Fatalf(msg string, v ...any) {
	logger.Fatalf(msg, v...)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Debugf(msg string, v ...any) {
	logger.Debugf(msg, v...)
}

func Warn(msg error) {
	logger.Warn(msg.Error())
}

func Warnf(msg string, v ...any) {
	logger.Warnf(msg, v...)
}

func Error(msg error) {
	logger.Error(msg.Error())
}

func Errorf(msg string, v ...any) {
	logger.Errorf(msg, v...)
}
