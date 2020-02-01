package main

import (
	"fmt"
	"test"
)

/**
工程项目管理
1.设置gopath file->settings->GO-GOPATH->project path
2. 项目必须包含一个main函数
3. 函数首字母大写为public,小写为private
4.调用不同包路径的需import包名并且用包名.函数名
*/
func main() {
	fmt.Println(test.Add(10, 20))
	fmt.Println(test.Minus(20, 10))
}
