package main

/**
* 这是第四个go concurrency 实例
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
