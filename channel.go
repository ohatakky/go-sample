package main

import "fmt"

func sender(ch chan string) {
	ch <- "data" // channelにデータを送信する
}

func main() {
	ch := make(chan string) // string型のデータを扱うchannel
	go sender(ch)           // goroutineにchannelを引数として渡す
	fmt.Println(<-ch)       // chからデータを受信する
}
