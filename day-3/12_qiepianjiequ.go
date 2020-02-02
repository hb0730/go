package main

import "fmt"

/**
切片截取
*/
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//[low,high,max]

	//array[0,len(array),len[array]]
	s1 := array[:]
	fmt.Println("s=", s1,
		"len = ", len(s1),
		"cap = ", cap(array))
	//操作数组一致
	s2 := array[1]
	fmt.Println("s2 = ", s2)

	s3 := array[0:4:5]
	fmt.Println("s3 = ", s3,
		"len = ", len(s3),
		"cap = ", cap(s3))

	s4 := array[:4]
	fmt.Println("s4 = ", s4,
		"len = ", len(s4),
		"cap = ", cap(s4))

	s5 := array[5:]
	fmt.Println("s5 = ", s5,
		"len = ", len(s5),
		"cap = ", cap(s5))

}
