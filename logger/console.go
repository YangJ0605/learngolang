package logger

import (
	"fmt"
	"os"
	"time"
)

type ConsoleLevel struct {
	level Level
}

// 构造函数
func NewConsoleLogger(levelStr string) *ConsoleLevel {

	consoleLogger := &ConsoleLevel{
		level: parseLogLevel(levelStr),
	}

	return consoleLogger
}

func (f *ConsoleLevel) log(l Level, format string, args ...interface{}) {
	if f.level > l {
		return
	}
	msg := fmt.Sprintf(format, args...)
	// 日志格式： [时间] [文件：行号] [函数名] [日志级别] 日志信息

	t := time.Now().Format("2006-01-02 15:04:05")
	fileName, line, funcName := GetCallerInfo(3)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s", t, fileName, line, funcName, getLevelStr(l), msg)
	fmt.Fprintln(os.Stdout, logMsg)

}

// 方法
func (c *ConsoleLevel) Debug(format string, args ...interface{}) {
	c.log(Debug, format, args...)
}

func (c *ConsoleLevel) Info(format string, args ...interface{}) {
	c.log(Info, format, args...)
}

func (c *ConsoleLevel) Warning(format string, args ...interface{}) {
	c.log(Warning, format, args...)
}

func (c *ConsoleLevel) Error(format string, args ...interface{}) {
	c.log(Error, format, args...)
}

func (c *ConsoleLevel) Fatal(format string, args ...interface{}) {
	c.log(Fatal, format, args...)
}

func (c *ConsoleLevel) Close() {}
