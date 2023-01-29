package main

import (
	"bytes"
	"fmt"
)

func main() {
	// n := 1
	// n1 := 2.1
	// n2 := n + n1 这样报错
	fmt.Printf("n2: %v\n", 1+2.1)
	const n = 1 // 未定类常量 没有声明类型的常量都叫做未定类常量
	const n1 = 2.1

	// fmt.Printf("a: %T\n", a)
	n2 := n + n1
	fmt.Printf("n2: %v\n", n2)

	// var a rune
	// var b byte
	var subnet string = "192.168.1.0/24"
	host := 100
	// var IP = "172.16.10.20"
	// var usage float32
	// usage = 23.8
	fmt.Print("网段" + subnet + "下有" + fmt.Sprint(host) + "台活跃主机。\n")

	var vendor1 bytes.Buffer
	vendor1.WriteString("Css")
	fmt.Printf("vendor1.String(): %v\n", vendor1.String())

	vendor2 := "Cisco"
	// 在Go中，索引返回的值为byte（byte是uint8的别名，uint8的中文叫做无符号8位整数)
	fmt.Println(vendor2[0]) // 打印的是 字母C对应的uint8的值
	fmt.Printf("vendor2[0]: %c\n", vendor2[0])
	fmt.Printf("vendor2[0]: %v\n", string(vendor2[0]))

	c := "中国人民aaaaaa"
	fmt.Printf("c: %v\n", c[:4])
	fmt.Printf("c: %v\n", string(([]rune(c))[:4]))
	// rune int32
	// 声明一个字符默认是rune
	d := '中'
	fmt.Printf("d: %v, type: %T, s: %c\n", d, d, d)

	var arr = [10]int{}
	fmt.Printf("arr的长度: %v, 容量: %v\n", len(arr), cap(arr))
	s1 := arr[:0]
	fmt.Printf("s1的长度: %v, 容量: %v\n", len(s1), cap(s1))

	s2 := arr[5:]
	fmt.Printf("s2的长度: %v, 容量: %v\n", len(s2), cap(s2))
}
