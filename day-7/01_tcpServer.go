package main

import (
	"fmt"
	"net"
)

/**
tcp服务器
*/
func main() {
	serverTest()
}

func serverTest() {
	//监听
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	// 阻塞等待用户连接
	conn, err1 := listen.Accept()
	defer conn.Close()
	if err1 != nil {
		panic(err1)
	}
	by := make([]byte, 1024)
	var n1 int
	n1, err1 = conn.Read(by)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("buf =", string(by[:n1]))
}
