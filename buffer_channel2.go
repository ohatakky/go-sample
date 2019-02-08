package main

import "fmt"

func worker(receiver chan int) {
	buffer := make(chan string, 10) // buffered channel size : 10
	go func() {
		for i := 0; i < 100; i++ {
			buffer <- "use" // データを送信するたび1つbufferする
			go func(i int) {
				receiver <- i
				<-buffer // データを送信したらbufferを1つ解放する
			}(i)
		}
	}()
}

func main() {
	receiver := make(chan int)
	worker(receiver)
	for i := 0; i < 100; i++ {
		fmt.Println(<-receiver)
	}
}
