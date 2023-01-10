package main

import "fmt"

// go的运算符

func main() {
	a := 10
	b := 20

	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(5 / 2): %v\n", (5 / 2)) // 2

	fmt.Printf("(5 %% 2): %v\n", (5 % 2)) // 1

	a++
	fmt.Printf("a: %v\n", a)
	a--
	fmt.Printf("a: %v\n", a)

	// 逻辑运算符
	// && || !

	// 位运算符
	// & | ^ >> <<
}
