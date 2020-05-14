package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go hello1()
	go hello2()

	fmt.Println("main 进入阻塞状态, 等待wg组中子goroutine结束...")
	wg.Wait()
	fmt.Println("main 解除阻塞...")
}

func hello1() {
	for i:=0; i<5; i++ {
		fmt.Println("hello1 -----", i)
	}
	wg.Done()  // 给wg等待组中的counter值减1, 同 wg.Add(-1)
}

func hello2() {
	defer wg.Done()
	for j:=0; j<5; j++ {
		fmt.Println("hello2 ==========", j)
	}
}
