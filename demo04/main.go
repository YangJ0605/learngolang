package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	n := 10
	fmt.Printf("10进制: %d\n", n)
	fmt.Printf("2进制: %b\n", n)
	fmt.Printf("8进制: %o\n", n)
	fmt.Printf("16进制: %X\n", n)

	m := 075
	x := 0xff

	fmt.Printf("m8进制: %o\n", m)
	fmt.Printf("x16进制: %X\n", x)
	fmt.Printf("x10进制: %d\n", x)

	fm := math.MaxFloat32
	fm2 := math.MaxFloat64

	fmt.Printf("fm: %v\n", fm)
	fmt.Printf("fm2: %v\n", fm2)

	var b1 bool   // 默认值为false
	var n1 int    // 默认值为0
	var s1 string // 默认值空字符串
	fmt.Printf("b1: %v %v %v\n", b1, n1, s1)

	// 转义
	fmt.Println("c:\\code\\golang")
	fmt.Println("\tc\nd")

	// 多行
	s2 := `
	1 
	2
	\t
	`
	fmt.Printf("s2: %v\n", s2)

	fmt.Printf("len(s2): %v\n", len(s2))

	s3 := "hello邹"
	fmt.Printf("len(s3): %v\n", len(s3))
	fmt.Printf("(s3 + s2): %v\n", (s3 + s2))
	fmt.Printf("len(s3 + s2): %v\n", len(s3+s2))
	s4 := fmt.Sprintf("%s - %s", s3, "云")
	fmt.Printf("s4: %v\n", s4)

	s5 := "ni hao sorry"
	fmt.Printf("s5 split: %v\n", strings.Split(s5, " "))
	fmt.Printf("strings.Split(s5, \" \"): %T\n", strings.Split(s5, " "))

	fmt.Printf("strings.Contains(s5, \"ni\"): %v\n", strings.Contains(s5, "ni11"))

	fmt.Printf("strings.HasPrefix(s5, \"ni\"): %v\n", strings.HasPrefix(s5, "ni11"))
	fmt.Printf("strings.HasSuffix(s5, \"ry\"): %v\n", strings.HasSuffix(s5, "y1"))

	fmt.Printf("strings.Index(s5, \"r\"): %v\n", strings.Index(s5, "r"))
	fmt.Printf("strings.LastIndex(s5, \"r\"): %v\n", strings.LastIndex(s5, "r"))

	fmt.Printf("strings.Join([]string{\"c\", \"c\"}, \"-\"): %v\n", strings.Join([]string{"c", "c"}, "-"))

	// byte uint8的别名
	// rune int32的别名

	var c byte = 'c'
	var d rune = 'c'
	fmt.Printf("c的值: %v c的类型%T\n", c, c)
	fmt.Printf("d: %T %v\n", d, d)

	s8 := "hello cc 邹"
	for i := 0; i < len(s8); i++ {
		fmt.Printf("s8[i]: %c\n", s8[i])
	}

	for _, v := range s8 {
		fmt.Printf("v: %c\n", v)
	}
}
