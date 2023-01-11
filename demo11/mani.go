package main

import "fmt"

func modify1(a int) {
	a = 1000
}

func modify2(a *int) {
	*a = 2000
}

func main() {
	a := 20
	b := &a

	a = 10

	fmt.Printf("b: %p, type: %T, value: %v\n", b, b, *b)
	fmt.Printf("a: %p\n", &a)

	modify1(a)
	fmt.Printf("a: %v\n", a)
	modify2(&a)
	fmt.Printf("a: %v\n", a)

	var kong *int // 声明指针类型 如果没有初始化会报错
	kong = new(int)
	*kong = 100 // 不初始化会报错
	fmt.Printf("kong: %p, value: %v\n", kong, *kong)

	ccc := new(map[string]string)
	fmt.Printf("ccc: %v\n", ccc)
}

// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
// new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
