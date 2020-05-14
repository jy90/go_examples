package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {
	fmt.Println("GOROOT -->", runtime.GOROOT())
	fmt.Println("os/platform -->", runtime.GOOS)
	fmt.Println("logical CPU nums -->", runtime.NumCPU())

	/*
	// runtime.Gosched()
	go func() {
		for i:=0; i<5; i++ {
			fmt.Println("goroutine...", i)
		}
	}()

	for i:=0; i<4; i++ {
		// 让出时间片,先让别的goroutine执行
		runtime.Gosched()
		fmt.Println("main ---", i)
	}
	*/


	// runtime.Goexit()
	go func() {
		fmt.Println("goroutine start...")
		hello()
		fmt.Println("goroutine stop...")
	}()
	time.Sleep(3 * time.Second)
}

func hello() {
	defer fmt.Println("hello defer !!!")
	// return  // 退出函数
	runtime.Goexit() // 退出goroutine
	fmt.Println("hello stop...")
}
