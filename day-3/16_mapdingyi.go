package main

import "fmt"

/**
map定义
*/
func main() {
	//test();
	//test2();
	//forMap();
	//deleteMap();
	test3()
}

//定义
func test() {
	var m1 = make(map[int]string, 2)
	fmt.Println("map = ", m1, "len = ", len(m1))
	m1[0] = "str1"
	m1[1] = "str2"
	m1[3] = "str3"
	fmt.Println("map =", m1, "len = ", len(m1))
}

// 初始化
func test2() {
	m1 := map[int]string{1: "str1", 2: "str2", 3: "str3"}
	fmt.Println(m1)
}

// map 遍历
func forMap() {
	m1 := map[int]string{1: "str1", 2: "str2", 3: "str3"}
	for key, value := range m1 {
		fmt.Println("key = ", key, "value = ", value)
	}
	//判断key是否存在
	// value为map的value，ok为map的key是否存在
	value, ok := m1[0]
	if ok {
		fmt.Println("value = ", value)
	} else {
		fmt.Println("key 不存在")
	}
}

// 删除
func deleteMap() {
	m1 := map[int]string{1: "str1", 2: "str2", 3: "str3"}
	delete(m1, 1)
	fmt.Println("m1 = ", m1)
}

/**
map做函数参数传递
实则为引用传递
*/
func test3() {
	m1 := map[int]string{1: "str1", 2: "str2", 3: "str3"}
	mapParams(m1)
	fmt.Println("m1 = ", m1)
}

func mapParams(m1 map[int]string) {
	delete(m1, 1)

}
