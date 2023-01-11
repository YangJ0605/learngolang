package main

import (
	"errors"
	"fmt"
)

// 函数 匿名函数 闭包

/*
	func 函数名(参数)(返回值){
    函数体
	}

*/

func sum(a, b int) int {
	return a + b
}

// 可变参数 可变参数在函数体中是切片类型
func sum2(x ...int) int {
	fmt.Printf("x: %v, type: %T\n", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 固定参数 与 可变参数， 可变参数一定要放到固定参数的后面
func sum3(x int, y ...int) int {
	res := x
	for _, v := range y {
		res += v
	}

	return res
}

// 多个返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

// 返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	// 如果return 为空 默认返回命名的返回值，如果return 后面有值就返回后面的值
	return 1, 2
}

// nil 可以作为切片的返回值

func foo(x string) []int {
	if x == "" {
		return nil
	}
	return []int{1, 2, 3}
}

// 全局变量 声名在函数外部 条件语句 循环语句外面的
// 局部变量 声明在函数 条件 循环语句里面的

// 定义函数类型
type calculation func(int, int) int

// 高阶函数
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calcOp(x, y int, op func(x, y int) int) int {
	return op(x, y)
}

// 函数作为返回值
func op(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		return nil, errors.New("无法识别的操作符")
	}
}

// 闭包和匿名函数

func adder() func(int) int {
	x := 0

	return func(i int) int {
		x += i
		return x
	}
}

func calc22(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

// panic recover
func funcA() {
	fmt.Println("funcA")
}

func funcB() {
	// recover 必须得在defer函数里面 并且必须在panic前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic B", err)
		}
	}()
	panic("funcB")

}

func funcC() {
	fmt.Println("funcC")
}

// defer
// 被defer的语句最后被执行，最后被defer的语句，最先被执行。
func deferFn() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}

func main() {
	num := sum(1, 3)
	fmt.Printf("num: %v\n", num)
	num2 := sum2(1, 4, 7, 8, 10)
	num3 := sum2(100, 10)

	fmt.Printf("num2: %v\n", num2)
	fmt.Printf("num3: %v\n", num3)

	num4 := sum3(1, 100, 99)
	fmt.Printf("num4: %v\n", num4)

	x, y := calc(5, 1)
	fmt.Printf("x: %v y: %v\n", x, y)

	x2, y2 := calc2(5, 1)
	fmt.Printf("x2: %v y2: %v\n", x2, y2)

	s := foo("")
	fmt.Printf("s: %v\n", s == nil)

	var c calculation
	c = sum
	fmt.Printf("c: %T\n", c)
	fmt.Printf("sum: %T\n", sum)
	cc := c(1, 2)

	fmt.Printf("cc: %v\n", cc)

	fmt.Printf("calcOp(1, 6, add): %v\n", calcOp(1, 6, add))

	addOp, error := op("")
	if error == nil {
		fmt.Printf("addOp(1, 344): %v\n", addOp(1, 344))
	} else {
		fmt.Printf("error: %v\n", error)
	}

	fn := adder()
	fmt.Printf("fn(10): %v\n", fn(10))
	fmt.Printf("fn(20): %v\n", fn(20))
	fmt.Printf("fn(30): %v\n", fn(30))

	fn2 := adder()
	fmt.Printf("fn2(100): %v\n", fn2(100))
	fmt.Printf("fn2(200): %v\n", fn2(200))

	f1, f2 := calc22(10)
	fmt.Printf("f1(1): %v f2(1): %v\n", f1(1), f2(1))
	fmt.Printf("f1(5): %v f2(2): %v\n", f1(5), f2(2))

	funcA()
	funcB()
	funcC()

	deferFn()

}
