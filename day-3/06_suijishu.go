package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
随机数
*/
func main() {
	//设置种子,只需一次
	// 如果参数为固定值,每次产生随机数相同
	// rand.Seed(666);
	rand.Seed(time.Now().UnixNano())
	//产生随机数
	for i := 0; i < 5; i++ {
		//随机数很大
		//fmt.Println("rand=", rand.Int());
		//随机数范围
		fmt.Println("rand = ", rand.Intn(100))
	}
}
