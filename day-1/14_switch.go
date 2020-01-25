package main

import "fmt"

/**
switch
*/
func main() {
	s1 := 10
	switch s1 {
	case 10:
		fmt.Printf("s1 =10")
		//break; //跳出switch语句
		//fallthrough; //不跳出switch语句
	case 20:
		fmt.Printf("s1 =20")
		//break;
		//fallthrough;
	default:
		fmt.Printf("s1 default")
	}
}
