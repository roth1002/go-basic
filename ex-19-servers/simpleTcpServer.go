package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func server() {
	server, err := net.Listen("tcp", ":6070")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleServerConnection(connection)
	}
}

func handleServerConnection(connect net.Conn) {
	var msg string
	err := gob.NewDecoder(connect).Decode(&msg) //傳參考去接
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	connect.Close()
}

func client() {
	for {
		client, err := net.Dial("tcp", "127.0.0.1:6070")
		if err != nil {
			fmt.Println(err)
			return
		}

		message := "Hello World"
		fmt.Println("Sending", message)
		err = gob.NewEncoder(client).Encode(message) //送出只需要送Value
		if err != nil {
			fmt.Println(err)
		}

		client.Close()
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
