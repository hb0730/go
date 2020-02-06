package main

import (
	"fmt"
	"net"
)

func main() {

	listen, err1 := net.Listen("tcp", ":8080")
	if err1 != nil {
		panic(err1)
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	by := make([]byte, 1024)
	n, er := conn.Read(by)
	if er != nil {
		panic(er)
	}
	fmt.Println("接收", string(by[:n]))
}
