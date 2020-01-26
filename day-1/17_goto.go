package main

import "fmt"

/**
goto语句
*/
func main() {
	fmt.Println("11111")
	//跳转
	goto end
	fmt.Println("22222")
end:
	fmt.Println("33333")
}
