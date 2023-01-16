package main

import (
	"fmt"
	"time"
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

	// ch1 := make(chan int, 100)
	// ch2 := make(chan int, 200)
	// go f1(ch1)

	// go f2(ch1, ch2)

	// for c := range ch2 {
	// 	fmt.Printf("c: %v\n", c)
	// }

	// test2()

	// selectDemo()
	demo2()
}

// 无缓冲通道补充
func test2() {

	ch := make(chan int)
	go recv(ch) // 创建一个goroutine从通道接收值
	// 这样做会直接报错 无缓冲的通道必须有至少一个接收方才能发送成功
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c <-chan int) {
	x := <-c
	fmt.Println("接受成功", x)
}

// select 多路复用
/*
	select {
	case <-ch1:
	//...
	case data := <-ch2:
	//...
	case ch3 <- 10:
	//...
	default:
	//默认操作
}

Select 语句具有以下特点。

可处理一个或多个 channel 的发送/接收操作。
如果多个 case 同时满足，select 会随机选择一个执行。
对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
*/

func selectDemo() {
	ch := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("x: %v\n", x)
		case ch <- i:
		}
	}
}

// demo2 通道误用导致的bug
func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		fmt.Println("===")
		return
	}
}
