package main

import (
	"fmt"
	"time"
)

func f1(ch chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("res: ", <-ch)
}

func send(ch chan int) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {

	// ch1 := make(chan bool)

	// go f1(ch1)
	// ch1 <- true // 无缓冲通道相当于同步通道 会一直等待别人接受它发送的值 不然会死锁
	// fmt.Println("main 结束")

	ch2 := make(chan int, 1)

	go send(ch2)

	for v := range ch2 {
		fmt.Printf("v: %v\n", v)
	}

}
