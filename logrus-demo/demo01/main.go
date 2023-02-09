package main

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func Demo01() {
	// 默认日志级别是info 所以debug级别的信息不出来
	fmt.Printf("logrus.GetLevel(): %v\n", logrus.GetLevel())
	// 更改日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 设置输出样式 text 或者json
	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	logrus.SetFormatter(&logrus.TextFormatter{
		// ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	// 添加额外信息
	logrus.WithField("project", "logrus-demo").Debug("cccc")
	myLog := logrus.WithFields(logrus.Fields{
		"name": "cc",
		"age":  18,
	})
	myLog.Info("info")
	myLog.Warn("warn")
	myLog.Error("error")
	myLog.Print("print")
	myLog.Info("info2")
}

func Demo02() {
	log := NewLog()
	log.Info("ddddddd")
	log.Error("ffffff")
}

func main() {
	Demo02()
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	if entry.HasCaller() {
		fmt.Println("开启caller")
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

func NewLog() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetLevel(logrus.DebugLevel)
	mLog.SetFormatter(&LogFormatter{})
	mLog.SetReportCaller(true)

	return mLog
}
