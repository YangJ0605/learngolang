package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go语言中的并发编程 goroutine

// 声明全局等待组变量
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello, goroutine", i)
	defer wg.Done() // 通知wg计数器减一
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A: ", i)
	}
	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B: ", i)
	}
	wg.Done()
}

func main() { // main函数开启一个主goroutine去执行main函数

	// 	for i := 0; i < 10000; i++ {
	// 		wg.Add(1)   // 登记一个goroutine
	// 		go hello(i) // 开启了一个goroutine去执行这个hello函数
	// 	}
	// 	fmt.Println("main goroutine dene")
	// 	// time.Sleep(time.Second * 1)
	// 	wg.Wait() // 阻塞 等待登记的goroutine完成

	runtime.GOMAXPROCS(4)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
