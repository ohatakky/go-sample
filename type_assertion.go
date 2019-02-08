package main

import (
	"fmt"
	"log"
)

func do(i interface{}) { // どんなinterfaceも受け付ける
	switch i.(type) {
	case int:
		ii := i.(int) // タイプアサーション
		ii *= 2
		fmt.Println(ii)
	case string:
		ss := i.(string)
		fmt.Println(ss)
	default:
		log.Fatal("error")
	}
}

func main() {
	do(10)
	do("pochi")
	do(true)
}
