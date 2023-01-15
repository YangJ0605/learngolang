package main

import (
	"fmt"
)

// channel 通道 队列的规则 先进先出

/*
	var 变量 chan 元素类型 引用类型
	 需要使用make函数初始化的就三种 chan 切片 map
*/

func test1() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int
	fmt.Printf("ch1: %v\n", ch1) // nil
	fmt.Printf("ch2: %v\n", ch2)
	fmt.Printf("ch3: %v\n", ch3)
	// 有缓冲通道 如果是无缓冲通道 不能存值 无缓冲又称为同步通道
	ch1 = make(chan int, 1) // 如果没有设置缓冲区大小 那就是无缓冲通道 无法发送值
	ch1 <- 10
	fmt.Printf("len(ch1): %v\n", len(ch1))
	// x := <-ch1
	fmt.Printf("ch1: %v\n", <-ch1)

	close(ch1)
}

// chan 有发送 接收 关闭三种操作
// chan <- 代表只能往通道发送值
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// ch1 <- chan int 代表只能从通道接收值
func f2(ch1 <-chan int, ch2 chan<- int) {
	for v := range ch1 {
		ch2 <- v * v
	}
	close(ch2)
}

// 单向通道
// chan <- 代表只能往通道发送值
// ch1 <- chan int 代表只能从通道接收值

func main() {
	// test1()
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)

	go f2(ch1, ch2)

	for c := range ch2 {
		fmt.Printf("c: %v\n", c)
	}
}
