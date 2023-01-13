package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func readFormFile() {
	// 打开文件
	file, err := os.Open("./main.go")

	if err != nil {
		fmt.Printf("open err: %v\n", err)
		return
	}
	fmt.Printf("file: %v\n", string(file.Name()))
	defer file.Close()

	var temp = make([]byte, 128)

	// 读取文件内容
	n, err := file.Read(temp)

	if err != nil {
		fmt.Printf("read err: %v\n", err)
		return
	}

	fmt.Printf("n: %v\n", n)
	fmt.Printf("temp: %v\n", string(temp))
}

func readAll() {
	file, err := os.Open("./1.txt")

	if err != nil {
		fmt.Printf("open err: %v\n", err)
		return
	}
	fmt.Printf("file: %v\n", string(file.Name()))
	defer file.Close()

	for {
		var temp = make([]byte, 128)

		// 读取文件内容
		n, err := file.Read(temp)

		if err == io.EOF {
			fmt.Println("读完了~")
			fmt.Printf("文件内容为: %v\n", string(temp[:n]))
			return
		}

		if err != nil {
			fmt.Println("read file err", err)
			return
		}
		fmt.Printf("读取到的字节数为: %d\n", n)
		fmt.Printf("文件内容为: %v\n", string(temp[:n]))
	}

}

// 通过buffer io 读取文件
func readByBufIo() {
	file, err := os.Open("./1.txt")

	if err != nil {
		fmt.Printf("open err: %v\n", err)
		return
	}
	fmt.Printf("file: %v\n", string(file.Name()))
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		s, err := reader.ReadString('\n')

		if err == io.EOF {
			if len(s) != 0 {
				fmt.Printf("最后一行文件内容: %v\n", s)
			}
			fmt.Println("文件读完了")

			return
		}

		if err != nil {
			fmt.Println("read file by buffer io failed, err: ", err)
			return
		}

		fmt.Printf("s: %v\n", s)
	}

}

func readByIoUtil() {
	content, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Printf("content: %v\n", string(content))
}

func main() {
	// readAll()
	// readByBufIo()
	readByIoUtil()
}
