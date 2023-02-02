package main

import (
	"logger"
)

func main() {
	fileLogger := logger.NewFileLogger("debug", "logs", "log.log")

	for {
		fileLogger.Info("这是一条info的日志%v", "2023.02.02")
		fileLogger.Error("这是一条Error的日志%v", "2023.02.02")
	}
	// file, err := os.OpenFile("logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err.Error())
	// }
	// fmt.Fprintln(file, "ffassadsadsadsadsadsad")
}
