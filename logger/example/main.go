package main

import (
	"logger"
)

var log logger.Logger

func main() {
	log = logger.NewConsoleLogger("debug")
	defer log.Close()

	// time.Sleep(5 * time.Second)
	log.Info("这是一条info信息")
	log.Error("这是一条error信息")

	// for {
	// 	log.Info("这是一条info信息")
	// 	log.Error("这是一条error信息")
	// }
}
