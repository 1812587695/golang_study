package main

/**
该主函数退出，goroutine来不及执行或者执行不完全
 */

import (
	"fmt"
)

func main() {
	// 通过make创建channel 类型为bool
	c := make(chan bool)

	go func() {
		fmt.Println("hello world!!")
		<-c // 取出channel这个变量的值，阻塞在这里
	}()
	close(c)
}
