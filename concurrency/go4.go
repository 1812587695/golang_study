package main

/**
* 这是第四个go concurrency 实例
  - 管道(Channel)是Go语言中比较重要的部分，经常在Go中的并发中使用
  - 管道是类型相关的，即一个管道只能传递一种类型的值。管道中的数据是先进先出的
	管道是什么？
	- 管道是Go语言在语言级别上提供的goroutine间的**通讯方式**，我们可以使用channel在多个goroutine之间传递消息。
	- channel是**进程内**的通讯方式，是不支持跨进程通信的，如果需要进程间通讯的话，可以使用Socket等网络方式。
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
	c <- true //给channel的变量c赋值true
}

// 等同于go3.go
