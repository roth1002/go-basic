package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for {
			chan1 <- "from tim"
			time.Sleep(time.Second * 1)
		}
	}() //很像Javascript吧XD

	go func() {
		for {
			chan2 <- "from randy"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-chan1:
				fmt.Println("message 1", msg1)
			case msg2 := <-chan2:
				fmt.Println("message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
