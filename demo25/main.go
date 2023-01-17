package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	file, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错, err:", err)
		return
	}
	name := "ffffff"
	fmt.Fprintf(file, "往文件中写入信息: %s\n", name)

	// Sprint 会返回一个新的字符串
	s1 := fmt.Sprint("沙河小王子")
	name = "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

	//Errorf
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Printf("w: %v\n", w)

	// 获取输入
	// testScan()

	// bufio
	bufioDemo()

}

func testScan() {
	var (
		name    string
		age     int
		married bool
	)
	// fmt.Scan(&name, &age, &married)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	// fmt.Printf("married: %T\n", married)
	fmt.Printf("扫描结果 name:%s age: %d, married: %t\n", name, age, married)
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Println("请输入内容:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("text: %#v\n", text)
}
