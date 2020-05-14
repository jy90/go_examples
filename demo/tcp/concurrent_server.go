package main

import (
	"net"
	"fmt"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端连接成功!!!")

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		fmt.Println(buf[:n])
		if "exit\n" == string(buf[:n]) {
			fmt.Println("客户端请求关闭连接...")
			return
		}

		if n == 0 {
			fmt.Println("连接已断开...")
			return
		}

		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Println("服务器读到的数据:", string(buf[:n]))
		// 小写转大写, 回发给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		go HandleConn(conn)
	}
}
