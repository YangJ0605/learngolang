package main

import "fmt"

// 空接口
type Any interface{}

type Dog struct{}

func main() {
	var x Any

	x = "你好"
	fmt.Printf("x: %v type: %T\n", x, x)

	x = 100
	fmt.Printf("x: %v type: %T\n", x, x)

	x = Dog{}
	fmt.Printf("x: %v type: %T\n", x, x)

	// 直接声明
	var x2 interface{} // nil
	fmt.Printf("x2: %v type: %T\n", x2, x2)

	Show(1)

	// 空接口作为map的值 可以保存任意值类型
	var studentInfo = make(map[string]interface{}, 0)
	studentInfo["name"] = "cc"
	studentInfo["age"] = 18
	fmt.Printf("studentInfo: %#v\n", studentInfo)
}

// 使用空接口函数可以传任何参数
func Show(a interface{}) {
	fmt.Printf("x: %v type: %T\n", a, a)
}
