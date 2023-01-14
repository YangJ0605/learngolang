package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 复制文件
func CopyFile(targetPath, srcPath string) (writer int64, err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("open 源文件 failed, err: %v\n", err)
		return
	}
	defer src.Close()

	target, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("open %s failed, err: %v\n", targetPath, err)
		return
	}
	defer target.Close()

	return io.Copy(target, src)
}

// 实现一个cat命令
func Cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')

		if err == io.EOF {
			fmt.Fprint(os.Stdout, "内容：%s", buf)
			break
		}

		fmt.Fprint(os.Stdout, "内容：%s", buf)
	}
}

func main() {
	// CopyFile("./3.txt", "./2.txt")
	flag.Parse() // 解析命令好参数

	if flag.NArg() == 0 {
		Cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))

		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err: %v\n", flag.Arg(i), err)
			continue
		}
		Cat(bufio.NewReader(f))
	}
}
