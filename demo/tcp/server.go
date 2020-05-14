package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端连接...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务器与客户端建立连接成功!!!")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	fmt.Println("服务器读到的数据是:", string(buf[:n]))

	conn.Write(buf[:n])
}