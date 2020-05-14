package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Are you ready?"))

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	fmt.Println("服务器回发:", string(buf[:n]))
}
