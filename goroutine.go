package main

import (
	"fmt"
	"time"
)

func goroutine1() {
	go func() {
		fmt.Println("goroutine1")
	}()
}
func goroutine2() {
	go func() {
		fmt.Println("goroutine2")
	}()
	time.Sleep(1 * time.Second)
}

func main() {
	goroutine1()
	goroutine2()
}
