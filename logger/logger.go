package logger

import "strings"

type Level uint16

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

const (
	Debug Level = iota
	Info
	Warning
	Error
	Fatal
)

func getLevelStr(l Level) string {
	switch l {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Debug"
	}
}

func parseLogLevel(l string) Level {
	l = strings.ToLower(l)
	switch l {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Debug
	}
}
