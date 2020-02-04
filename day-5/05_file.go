package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

/**
文件操作
*/
func main() {
	test1()
}

//写入
func test1() {
	var path string = "./test.txt"
	var srcPath string = "./test2.txt"
	WriteFile(path)
	ReadFile(path)
	ReadRowFile(path)
	copyFile(path, srcPath)
}

/**
不存在则创建,存在清空后写入
*/
func WriteFile(path string) {
	var file *os.File
	var err error
	if IsExist(path) {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	} else {
		file, err = os.Create(path)
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	intn := rand.Int()
	s := fmt.Sprintf("sss%d", intn)
	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}

}

/**
读取文件
*/
func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	n, err1 := file.Read(buffer)
	if err1 != nil && err1 != io.EOF {
		panic(err1)
	}
	fmt.Println(string(buffer[:n]))
}

/**
行读
*/
func ReadRowFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		by, err1 := r.ReadBytes('\n')
		fmt.Println("byte = ", string(by))
		if err1 != nil && err1 == io.EOF {
			break
		} else if err1 != nil {
			panic(err1)
		}
	}
}

/**
文件copy
*/
func copyFile(dtsPath string, srcPath string) {
	file, err := os.Open(dtsPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if IsExist(srcPath) {
		err = os.Remove(srcPath)
		if err != nil {
			panic(err)
		}
	}
	srcFile, err1 := os.Create(srcPath)
	if err1 != nil {
		panic(err1)
	}
	rbufer := make([]byte, 1024)
	for {
		n, err2 := file.Read(rbufer)
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
		}
		srcFile.Write(rbufer[:n])
	}
}

/**
文件是否存在
*/
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsExist(err) || err == nil
}
