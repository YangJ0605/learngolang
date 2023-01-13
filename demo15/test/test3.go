package main

import "fmt"

// 类型断言
func Test(x interface{}) {
	switch t := x.(type) {
	case string:
		fmt.Printf("t: %v is string\n", t)
	case bool:
		fmt.Printf("t: %v is bool\n", t)
	case int:
		fmt.Printf("t: %v is int\n", t)
	default:
		fmt.Println("不支持的类型")
	}
}

func main() {
	var x interface{}
	x = 100

	ret, ok := x.(int8)        // 猜错返回类型对应的零值
	fmt.Printf("ok: %v\n", ok) // false
	fmt.Printf("ret: %v\n", ret)

	Test(x)
}
