package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type LogMessage struct {
	message  string
	levelStr string
}

type FileLogger struct {
	level              Level
	fileName, filePath string
	file               *os.File
	errFile            *os.File
	maxSize            int64
	logMessageChan     chan *LogMessage
}

// 构造函数
func NewFileLogger(levelStr string, filePath, fileName string) *FileLogger {

	fileLogger := &FileLogger{
		level:          parseLogLevel(levelStr),
		fileName:       fileName,
		filePath:       filePath,
		maxSize:        10 * 1024 * 1024,
		logMessageChan: make(chan *LogMessage, 10000),
	}
	fileLogger.initFile()
	return fileLogger
}

func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("open log file %s file %v", logName, err))
	}
	f.file = fileObj

	errLogName := fmt.Sprintf("error_%s", f.fileName)
	errLogName = path.Join(path.Dir(logName), errLogName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("open error log file %s file %v", errLogName, err))

	}
	f.errFile = errFileObj
	go f.writeMessageFromChan()
}

func (f *FileLogger) checkShouldSplit(file *os.File) bool {
	// fmt.Printf("file.Name(): %v\n", file.Name())
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	// fmt.Printf("fileInfo: %v\n", fileInfo)
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

func (f *FileLogger) splitLogFile(file *os.File) *os.File {

	fileName := file.Name()
	backupName := fmt.Sprintf("%s_%s.backup", fileName, time.Now().Format("2006_01_02_15_04_05_000"))
	file.Close()
	err := os.Rename(fileName, backupName)
	if err != nil {
		fmt.Printf("rename file err: %v\n", err.Error())
	}
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(fmt.Errorf("open Log file %s file %v", fileName, err))
	}

	return fileObj

}

func (f *FileLogger) writeMessageFromChan() {
	for logMessage := range f.logMessageChan {
		// 检查文件日志是否超过了maxSize
		if f.checkShouldSplit(f.file) {
			f.file = f.splitLogFile(f.file)
		}

		fmt.Fprintln(f.file, logMessage.message)

		l := parseLogLevel(logMessage.levelStr)

		if l >= Error {
			if f.checkShouldSplit(f.errFile) {
				f.errFile = f.splitLogFile(f.errFile)
			}
			fmt.Fprintln(f.errFile, logMessage.message)
		}
	}
}

func (f *FileLogger) log(l Level, format string, args ...interface{}) {
	if f.level > l {
		return
	}
	msg := fmt.Sprintf(format, args...)
	// 日志格式： [时间] [文件：行号] [函数名] [日志级别] 日志信息
	t := time.Now().Format("2006-01-02 15:04:05")
	fileName, line, funcName := GetCallerInfo(3)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s", t, fileName, line, funcName, getLevelStr(l), msg)

	logMessage := &LogMessage{
		message:  logMsg,
		levelStr: getLevelStr(l),
	}

	select {
	case f.logMessageChan <- logMessage:
	default:
		<-f.logMessageChan             // 丢失第一条数据
		f.logMessageChan <- logMessage // 写入最新的数据
	}

}

// 方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(Debug, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(Info, format, args...)
}

func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(Warning, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(Error, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(Fatal, format, args...)
}

func (f *FileLogger) Close() {
	f.errFile.Close()
	f.file.Close()
}
