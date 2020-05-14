package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	/*
	ch := make(chan int)

	fmt.Printf("%T, %v\n", ch, ch)

	go func() {
		for i:=0; i<5; i++ {
			fmt.Println("子goroutine ---->", i)
		}
		ch <- 10
	}()

	data := <- ch
	fmt.Println(data)
	*/

	/*
	ch1 := make(chan int)

	go sendData(ch1)

	for v := range ch1 {
		fmt.Println("读取数据 --->", v)
	}
	*/

	fmt.Println("=======================")
	ch2 := make(chan string, 4) // 带缓冲区的通道
	go sendData2(ch2)
	for {
		v, ok := <- ch2
		if !ok {
			fmt.Println("读完了...", ok)
			break
		}
		fmt.Println("\t读取的数据是:", v)
	}

	fmt.Println("主程序结束 !!!")
}

func sendData(ch chan int) {
	for i:=0; i<10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
	close(ch) // 通知对方, 关闭通道
}

func sendData2(ch chan string) {
	for i:=0; i<10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine正在写入第 %d 个数据\n", i)
	}
	close(ch)
}


ch01 :=  make(chan string) // 双向通道, 可读写
ch02 := make(<-chan string)  // 只读通道
ch03 := make(chan<- string)  // 只写通道
