package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
冒泡排序
*/
func main() {
	var array [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println(array)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
