package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "bob"}
}

func (e *UserNotFound) Error() string { // エラーでは、ポインタとして例ー部することが推奨されている。
	return fmt.Sprintf("%v not Found.", e.Username)
}

func main() {

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
