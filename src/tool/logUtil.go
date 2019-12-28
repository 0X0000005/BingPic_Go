package tool

type ILogger interface {
	Debug(format string,output ...interface{})
	Info(format string,output ...interface{})
	Warn(format string,output ...interface{})
	Error(format string,output ...interface{})
	Fatal(format string,output ...interface{})
	Panic(format string,output ...interface{})
}