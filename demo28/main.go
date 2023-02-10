package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("hello world")

	io.Copy(os.Stdout, r)
}

func testCopy2() {
	r := strings.NewReader("hello world")

	buf := make([]byte, 1)

	io.CopyBuffer(os.Stdout, r, buf)
}

func testIoutil() {
	r := strings.NewReader("cccccccc")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	fmt.Printf("string(b): %v\n", string(b))
}

func testBufIo() {
	r := strings.NewReader("hello world \n cccc")

	r2 := bufio.NewReader(r)

	s, _ := r2.ReadString(' ') // 以什么结尾
	fmt.Printf("s: %v\n", s)
}

func main() {
	// testCopy2()
	// testIoutil()
	testBufIo()
}
