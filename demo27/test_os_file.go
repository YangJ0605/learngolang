package main

import (
	"fmt"
	"io"
	"os"
)

func openClose() {
	// 只读
	// f, err := os.Open("4.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Printf("f.Name(): %v\n", f.Name())
	// f.Close()

	f, err := os.OpenFile("4.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
}

func ceateFile() {
	// 创建临时文件
	f, err := os.CreateTemp("", "temp")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func readFile() {
	f, _ := os.Open("main.go")

	for {
		buf := make([]byte, 3)
		n, err := f.Read(buf)
		fmt.Printf("string(buf): %v\n", string(buf))
		fmt.Printf("n: %v\n", n)
		if err == io.EOF {
			break
		}
	}
	f.Close()
}

func readDir() {
	de, _ := os.ReadDir("a")

	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}
func writeFile() {
	f, _ := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	f.Write([]byte("cccc"))
	f.WriteString("ffffff")
	f.Close()
}

func main() {
	// openClose()
	// ceateFile()
	// readFile()
	// readDir()
	writeFile()
}
