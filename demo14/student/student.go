package student

import "fmt"

// 函数和变量的大小写意味着对其他地方是否可以用
// 大写 可以访问 小写不行

var CommonConfig = []string{"cc"}
var a = 0

type Student struct {
	Name string
	Age  int8
}

func Make() {
	fmt.Println("Make in student")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("Add(1, 2): %v\n", Add(1, 2))
}

func init() {
	fmt.Println("int run in student")
}
