package main

/**
* 这是第6个go concurrency 实例
 */

import (
	"fmt"
)

func main() {
	// 通过make创建channel 类型为bool,设置缓存
	c := make(chan bool, 1)

	go func() {
		fmt.Println("hello world!!")

		c <- true
		// 如果这里不关闭就有报错 fatal error: all goroutines are asleep - deadlock
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
