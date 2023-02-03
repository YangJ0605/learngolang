package main

import (
	"logger"
)

var log logger.Logger

func main() {
	log = logger.NewFileLogger("debug", "./", "logs.log")
	defer log.Close()

	for {
		log.Info("这是一条info信息")
		log.Error("这是一条error信息")
	}

}
