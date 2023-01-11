package main

import "fmt"

// 类型别名和自定义类型

// 自定义类型
type MyInt int

// 类型别名
type Int2 = int

func main() {
	var a MyInt
	var b Int2
	a = 1
	b = 2

	fmt.Printf("a: %v, type: %T\n", a, a)
	fmt.Printf("b: %v, type: %T\n", b, b)
}
