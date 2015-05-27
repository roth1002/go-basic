package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Leo登出")
	fmt.Println(err)
}
