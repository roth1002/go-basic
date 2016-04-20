package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

func Command(ws *websocket.Conn) {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(ws)

	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		_, err = w.Write(data)
		if err != nil {
			panic(err)
		}
		w.Flush()

		fmt.Printf("您:\t> ")
	}
}

func main() {
	fmt.Printf("Welcome chatroom!")

	origin := "http://127.0.0.1:5000/"
	url := "ws://127.0.0.1:5000/chatroom"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	r := bufio.NewReader(ws)
	//w := bufio.NewWriter(os.Stdout)

	go Command(ws)

	for {
		data, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Printf("disconnected\n")
			os.Exit(0)
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("\r%sMy\t> ", data)
	}
}
