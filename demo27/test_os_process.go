package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 父亲id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())
	// 环境变量
	// fmt.Printf("os.Environ(): %v\n", os.Environ())
	// 获取某个环境变量
	fmt.Printf("os.Getenv(\"GOPATH\"): %v\n", os.Getenv("GOPATH"))
	s, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Printf("s: %v\n", s)
	}
}
