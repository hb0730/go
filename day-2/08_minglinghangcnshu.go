package main

import (
	"fmt"
	"os"
)

/**
获取命令行参数
*/
func main() {
	//接收用户参数
	strings := os.Args
	fmt.Printf("os args length %d", len(strings))
	for i, e := range strings {
		fmt.Printf("os args[%d]=%s", i, e)
	}
}
