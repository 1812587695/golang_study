package main

/**
* 这是第一个go concurrency 实例
  为什么这里什么都不输出
*/

import (
	"fmt"
)

// main是主程序，所有生命周期是main为主，不会管其他进程
// go Hello() 调用go启动一个进程运行，这个时候main函数没得等待Hello函数这个进程运行完就自动结束了
func main() {
	go Hello() // 调用下面的函数，不会打印
}

func Hello() {
	fmt.Println("hello wrod!!!") // 不会输出
}

// 如何解决这个问题了，请看下面分解：go2.go
