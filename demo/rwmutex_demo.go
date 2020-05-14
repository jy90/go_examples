package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	// 多个同时读取
	// wg.Add(2)
	// go readData(1)
	// go readData(2)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
}

func readData(i int) {
	defer wg.Done()
	fmt.Println("read start...", i)

	rwMutex.RLock()  // 读操作上锁
	fmt.Println("reading...", i)
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()  // 读操作解锁
	fmt.Println("read over...", i)
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println("write start... --->", i)

	rwMutex.Lock()  // 写操作上锁
	fmt.Println("writing... --->", i)
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()  // 写操作解锁
	fmt.Println("write over... --->", i)
}
