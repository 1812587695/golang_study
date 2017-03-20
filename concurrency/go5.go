package main

/**
* 这是第5个go concurrency 实例
 */

import (
	"fmt"
)

func main() {
	// 通过make创建channel 类型为bool,设置缓存
	c := make(chan bool, 1)

	// 这里会输出
	//	go func() {
	//		fmt.Println("hello world!!")
	//		c <- true //有缓存必须先存进去，后取出来
	//	}()
	//	<-c // 后取出来就阻塞等待，如果是下面这种就不可以了，因为有缓存是异步的

	// 如何没有设置了有缓存这里就的阻塞的，设置缓存了就不会管了，程序直接运行完毕
	// 结论有缓存是异步的，无缓存是同步阻塞的
	// 下面不会输出
	go func() {
		fmt.Println("hello world!!")
		<-c

	}()
	c <- true

}
