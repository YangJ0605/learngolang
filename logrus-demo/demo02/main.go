package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

type MyHook struct {
	Write *os.File
}

func (hook *MyHook) Levels() []logrus.Level {
	// 设置哪些级别的日志才会生效
	// return logrus.AllLevels
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	s, _ := entry.String()
	hook.Write.Write([]byte(s))
	entry.Data["name"] = "cc"
	return nil
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		// ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		// FullTimestamp:   true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	file, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.AddHook(&MyHook{
		Write: file,
	})
	logrus.Debug("ffff")
	logrus.Error("error")
}
