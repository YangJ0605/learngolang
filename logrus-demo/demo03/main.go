package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// logrus formatter
type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var fileName string
	var len int
	var msg string
	if entry.HasCaller() {
		fileName = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
		msg = fmt.Sprintf("[%s] %s [%s:%d] %s\n", strings.ToUpper(entry.Level.String()), timestamp, fileName, len, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s] %s  %s\n", strings.ToUpper(entry.Level.String()), timestamp, entry.Message)
	}

	return []byte(msg), nil
}

func main() {
	// logrus.SetFormatter(&LogFormatter{})
	// err := os.MkdirAll(path.Join("logs", "2023-02-010", "cc", "dd"), 0666)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Printf("time.Now().Local().Format(\"2006-01-02 15:04:05\"): %v\n", time.Now().Local().Format("2006-01-02 15:04:05"))

	// openFile只能在当前已经存在的目录打开或者创建文件，不能创建目录

	// logFileWriter, err := InitFileLog("./", "test")

	// logrus.Error("cccc")

	// if err != nil {
	// 	return
	// }

	// defer func() {
	// 	if logFileWriter.file != nil {
	// 		logFileWriter.file.Close()
	// 	}
	// }()

	// InitHookLog("logs/cc", "appDemo")

	// for {
	// 	logrus.Error("error")
	// 	time.Sleep(20 * time.Second)
	// 	logrus.Warn("warning")
	// }

	// InitLevel("logrus_study/log_level")
	// logrus.Errorln("你好")
	// logrus.Errorln("err")
	// logrus.Warnln("warn")
	// logrus.Infof("info")
	// logrus.Println("print")
}
