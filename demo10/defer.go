package main

import "fmt"

// defer 案例
func foo1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5 x是在函数体里面声明的，所以defer不会影响到返回值
}

func foo2() (x int) {
	defer func() {
		fmt.Printf("x: %v\n", x)
		x++
		fmt.Printf("x: %v\n", x)
	}()
	return 5 // 6 return 的时候把 5 赋值给了x，然后x++，导致x变成6
}

func foo3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}
func foo4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}

// 如果函数返回值中声明了变量，那么在defer中对其操作是可以影响到返回值的，如果没有，则不会

func main() {
	fmt.Println(foo1()) // 5
	fmt.Println(foo2()) // 6
	fmt.Println(foo3()) // 5
	fmt.Println(foo4()) // 5

	fmt.Printf("c(): %v\n", *c())
}
