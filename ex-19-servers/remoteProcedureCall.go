package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))
	server, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(connection)
	}
}

func client() {
	connection, err := rpc.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = connection.Call("Server.Negate", int64(9999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Nagete(9999) =", result)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
