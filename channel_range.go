package main

import "fmt"

func goroutine(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //使い終わったchannelはcloseする。
}

func main() {
	ch := make(chan int)
	go goroutine(ch)

	for v := range ch { // nilなchannelから読み込むとerrorになるのでclose()が必要
		fmt.Println(v)
	}
}
