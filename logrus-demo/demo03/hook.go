package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type SplitTimeHook struct {
	file                       *os.File
	logPath, fileTime, appName string
}

func (s *SplitTimeHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (s *SplitTimeHook) Fire(entry *logrus.Entry) error {
	time := entry.Time.Format("2006-01-02_15_04")

	str, _ := entry.String()

	if s.fileTime == time {
		fmt.Println("直接写入")
		s.file.Write([]byte(str))
		return nil
	}

	s.file.Close()

	err := os.MkdirAll(path.Join(s.logPath, time), 0666)

	if err != nil {
		logrus.Error(err)
		return err
	}

	fileName := path.Join(s.logPath, time, fmt.Sprintf("%s.log", s.appName))
	fmt.Printf("fileName: %v\n", fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	s.file = file
	s.fileTime = time

	s.file.Write([]byte(str))

	return nil
}

func InitHookLog(logPath, appName string) {
	fileTime := time.Now().Local().Format("2006-01-02_15_04")

	err := os.MkdirAll(path.Join(logPath, fileTime), 0666)

	if err != nil {
		logrus.Error(err.Error())
		return
	}
	fileName := path.Join(logPath, fileTime, fmt.Sprintf("%s.log", appName))

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	fileHook := SplitTimeHook{file: file, logPath: logPath, fileTime: fileTime, appName: appName}

	logrus.AddHook(&fileHook)

}
