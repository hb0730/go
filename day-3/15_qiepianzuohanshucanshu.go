package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
切片做函数参数（引用类型）
*/
func main() {
	array := make([]int, 4)
	initData(array)
	fmt.Println(array)
	modify(array)
	fmt.Println(array)
}
func initData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}
func modify(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
