package main

import (
	"fmt"
	"time"
)

func main() {
	// 临界资源
	a := 1

	go func() {
		a = 2
		fmt.Println("goroutine ---", a)
	}()

	a = 3
	time.Sleep(1)
	fmt.Println("main ---", a)
}

// go run -race demo_race.go
