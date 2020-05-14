package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var ticket = 10

var (
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(4)

	go sellTicket("售票口1")
	go sellTicket("售票口2")
	go sellTicket("售票口3")
	go sellTicket("售票口4")

	wg.Wait()
	fmt.Println("程序结束...")
}

func sellTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println("售出:", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println("售罄, 没票了...")
			break
		}
		mutex.Unlock()
	}
}
