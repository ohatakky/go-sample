package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	wg.Done() // goroutine1回終了
}

func task2(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	wg.Done() // goroutine1回終了
}

func Concurrency() {
	var wg sync.WaitGroup // 宣言
	wg.Add(2)             // goroutineが2回終了するまで、
	go task1(&wg)
	go task2(&wg)
	wg.Wait() // 待つ
	fmt.Println("done")
}
