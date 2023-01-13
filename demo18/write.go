package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入操作
func writeFile() {
	file, err := os.OpenFile("./2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}

	str := "\nccccc\n"

	file.Write([]byte(str))
	file.WriteString("hello cccccccc")

	defer file.Close()
}

func writeFileByBufIo() {
	file, err := os.OpenFile("./2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("ffffff") // 将内容写入缓冲区

	// 将缓存区的内容写入磁盘
	writer.Flush()
}

func writeByIoUtil() {
	str := `
		fffff
		fffffff
		dddd
		ffff
	`
	err := ioutil.WriteFile("./2.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("write file failed, err", err)
		return
	}
}

func main() {
	// writeFile()
	// writeFileByBufIo()
	writeByIoUtil()
}
