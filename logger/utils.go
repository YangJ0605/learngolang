package logger

import (
	"path"
	"runtime"
)

func GetCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileFullPath, line, ok := runtime.Caller(skip)

	if !ok {
		return
	}

	fileName = path.Base(fileFullPath)

	funcName = runtime.FuncForPC(pc).Name()

	return
}
