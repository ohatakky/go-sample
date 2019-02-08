package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}
type Dog struct {
	Name string
}

func (p *Person) Say() string { // 構造体Personは、Sayメソッドを実装したことにより、Humanインターフェースを満たした。
	return p.Name
}

func Drive(human Human) {
	if human.Say() == "yoshida" {
		fmt.Println("run")
	} else {
		fmt.Println("get out")
	}
}

func main() {
	var yoshida Human = &Person{"yoshida"}
	Drive(yoshida)

	var okamura Human = &Person{"okamura"}
	Drive(okamura)

	// var dog Dog = Dog{"pochi"}
	// Drive(&dog) // Say()メソッドを実装していない、つまりHumanインターフェースを満たしていないのでDriveを使おうとするとエラーになる。
}
