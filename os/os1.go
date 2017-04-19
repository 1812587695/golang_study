package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录是:", dir) //当前的目录是: D:\test 我的环境是windows 如果linix 就是/xxx/xxx

	// 获取环境变量
	path := os.Getenv("GOROOT")
	fmt.Println("环境变量GOPATH的值是:", path)
}
