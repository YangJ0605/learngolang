package main

import (
	"bufio"
	"fmt"
	"net"
)

// 网络编程
// go语言实现TCP通信

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte("ok-" + recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listen.Close()

	fmt.Println("server run at 127.0.0.1:8080")

	for {
		conn, err := listen.Accept() // 建立链接
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}
}
