package main

import "net"

/**
tcp客户端
*/
func main() {
	clientTest()
}
func clientTest() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("are you ok"))
}
