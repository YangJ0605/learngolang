package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

// 自定义writer
type LogFileWriter struct {
	file                       *os.File
	logPath, fileDate, appName string
}

func (lw *LogFileWriter) Write(data []byte) (n int, err error) {
	if lw == nil {
		return 0, errors.New("LogFileWriter is nil")
	}

	if lw.file == nil {
		return 0, errors.New("LogFileWriter.file is nil")
	}

	// 获取日期跟之前结构体初始化的日期做对比
	fileDate := time.Now().Local().Format("2006-01-02")

	// 当最新日期和结构体的日期不对的时候，需要重新创建文件夹
	if lw.fileDate != fileDate {
		lw.file.Close()
		// 创建文件夹
		err = os.MkdirAll(path.Join(lw.logPath, fileDate), 0666)

		if err != nil {
			return 0, err
		}

		fileName := path.Join(lw.logPath, fileDate, fmt.Sprintf("%s-%s.log", lw.appName, fileDate))

		lw.fileDate = fileDate
		lw.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			return 0, err
		}

	}
	n, e := lw.file.Write(data)

	return n, e

}

func InitFileLog(logPath, appName string) (*LogFileWriter, error) {
	fileDate := time.Now().Local().Format("2006-01-02")

	// 创建文件夹
	err := os.MkdirAll(path.Join(logPath, fileDate), 0666)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// 拼接文件地址
	fileName := path.Join(logPath, fileDate, fmt.Sprintf("%s-%s.log", appName, fileDate))

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	fileWriter := LogFileWriter{file: file, fileDate: fileDate, logPath: logPath, appName: appName}

	// logrus.SetOutput(os.Stdout)

	writers := []io.Writer{
		os.Stdout,
		&fileWriter,
	}

	logrus.SetOutput(io.MultiWriter(writers...))
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&LogFormatter{})
	return &fileWriter, nil
}
