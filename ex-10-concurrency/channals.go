package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)
	go tim(c) //相當於發起thread
	go randy(c)
	go receiveChannalAndPrint(c)
	var input string
	fmt.Scanln(&input) //故意讓主程式停住
}

func tim(tim chan<- string) { // 傳入channal 然後一直將 "tim" send進去 tim這個channal
	for i := 1; ; i++ { // go's while ( true ) in C lang
		tim <- "tim"
	}
}

func receiveChannalAndPrint(leo <-chan string) { // 傳入 channal 然後每秒接收channal的內容
	for { // another go's while ( true ) in C lang
		msg := <-leo
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func randy(randy chan<- string) { //可以將傳入方向放在argument宣告  方便讀懂code  如果是雙向都要的時候就不指定方向
	for i := 0; ; i++ {
		randy <- "randy"
	}
}
