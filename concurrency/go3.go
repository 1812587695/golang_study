package main

/**
* 这是第三个go concurrency 实例
  解决上go2.go的疑问
  引入：channel，通过通信来共享内存
	- channel是goroutime沟通的桥梁，大都是阻塞同步的
	- 通过make创建，clone关闭
	- 可以使用for range 来迭代不断操作channel
	- 可以设置单向或者双向通道
	- 可以设置缓存大小，在未背填满前不会发生阻塞
*/

import (
	"fmt"
)

func main() {
	// 通过make创建channel 类型为bool
	c := make(chan bool)

	go func() {
		fmt.Println("hello world!!")
		c <- true //给channel的变量c赋值true
	}()
	<-c // 取出channel这个变量的值,阻塞在这里
}

// 等同于go4.go
