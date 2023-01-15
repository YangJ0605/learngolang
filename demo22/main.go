package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发安全和锁

var (
	x  int
	wg sync.WaitGroup
	// 互斥锁
	lock sync.Mutex
	// 读写锁
	rwLock sync.RWMutex
)

// 读写锁的效率远远高于互斥锁 适合读多写小的场景

// 互斥锁
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁 x唯一
		x += 1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

// 读写互斥锁
func read() {
	// lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	// lock.Unlock()
	wg.Done()
}

func write() {
	rwLock.Lock()
	// lock.Lock()
	x += 1
	time.Sleep(time.Microsecond * 10)
	rwLock.Unlock()
	// lock.Unlock()
	wg.Done()
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Printf("x: %v\n", x)

	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()

	fmt.Println(time.Now().Sub(start))
}
