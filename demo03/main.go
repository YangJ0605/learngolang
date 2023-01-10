package main

import (
	"fmt"
)

// 常量声明
const pi = 3.1415
const e = 2.7

// 全局变量 声明不赋值 默认为false
var isOk bool

// 所有的值都一样
const (
	n1 = "CC"
	n2
	n3
)

// iota从0开始只能在常量中出现，每增加一行就加一
const (
	n4 = iota
	n5 = iota
	n6
	_
	n7
)

// 当const出现的时候 iota就重置为0
const (
	n8 = iota
	n9
)

const (
	n10, n11 = iota + 1, iota + 2
	n12, n13
	n14, n15
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB // 1 << (10 * iota) 后面都是这样 iota会依次累加
	GB
	TB
	PB
)

func main() {
	// 局部变量声明并且赋值
	var a string = "cc"

	// pi = 123
	// 短变量声明
	name := "cc"

	fmt.Printf("isOk: %v\n", isOk)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("pi: %v\n", pi)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("n1: %v\n", n1)
	fmt.Printf("n2: %v\n", n2)
	fmt.Printf("n3: %v\n", n3)
	fmt.Printf("n4: %v\n", n4)
	fmt.Printf("n5: %v\n", n5)
	fmt.Printf("n6: %v\n", n6)
	fmt.Printf("n7: %v\n", n7)
	fmt.Printf("n8: %v\n", n8)
	fmt.Printf("n9: %v\n", n9)
	fmt.Println("nnn", n10, n11, n12, n13, n14, n15)
	fmt.Println("KB", KB, MB, GB, TB, PB)
}
