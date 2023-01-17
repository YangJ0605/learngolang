package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client demo

func main() {
	// 与服务端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}

	fmt.Println("connect server success")

	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 给服务端发送消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("send data err:", err)
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("read failed, err: ", err)
			return
		}

		fmt.Println("收到server发送的数据为:", string(buf[:n]))

	}
}
