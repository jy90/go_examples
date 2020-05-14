package main

import (
	"fmt"
	"net"
	"time"
	// "strings"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("net.ListenUDP error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("udp服务器通信socket创建完成.")

	buf := make([]byte, 4096)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP error:", err)
			continue
		}
		fmt.Printf("从客户端 %v 读取到的数据是 %s\n", clientAddr, string(buf[:n]))

		go func() {
			_, err = conn.WriteToUDP([]byte(time.Now().String()+"\n"), clientAddr)
			if err != nil {
				fmt.Println("conn.WriteToUDP error:", err)
				return
			}
		}()
	}
}
