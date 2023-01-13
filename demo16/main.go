package main

import (
	"errors"
	"fmt"
)

// Error接口和错误处理

// Go 语言中把错误当成一种特殊的值来处理，不支持其他语言中使用try/catch捕获异常的方式。

// type error interface {
// 	Error() string
// }

func Test1(id int) (interface{}, error) {
	if id%2 == 0 {
		return "ccc", nil
	}
	return nil, errors.New("错误id")
}

func main() {
	value, err := Test1(1)
	if err != nil {
		fmt.Printf("有错误，err: %v", err)
		return
	}
	fmt.Printf("value: %v\n", value)
}
