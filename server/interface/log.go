package interfaces

type Logger interface {
	Info(msg string)
	Infof(msg string, v ...any)
	Debug(msg string)
	Debugf(msg string, v ...any)
	Warn(msg string)
	Warnf(msg string, v ...any)
	Fatal(msg string)
	Fatalf(msg string, v ...any)
	Error(msg string)
	Errorf(msg string, v ...any)
}
