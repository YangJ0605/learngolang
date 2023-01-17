package main

import (
	"fmt"
	"log"
	"os"
)

// log 包的使用

func main() {
	log.Println("日志log。。。。")
	v := "ccc"
	log.Printf("name: %v\n", v)
	// log.Fatalln("触发fatal的日志")
	// log.Panicln("触发panic的日志")

	// test1()
	test2()
}

func test1() {
	// 配置log
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("【order模块】")
	log.Println("这是一条很普通的日志222。")
}

func test2() {
	logFile, err := os.OpenFile("./a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("【order模块】")
	log.Println("这是一条很普通的日志222。")
}
