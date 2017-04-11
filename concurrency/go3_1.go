package main

/**
* 这是第三个go concurrency 实例

 */

import (
	"fmt"
)

func main() {
	// 通过make创建channel 类型为bool
	c := make(chan bool)

	go func() {
		fmt.Println("hello world!!")
		close(c) //在里面给ch赋值后（或者close（ch））后,才能继续往后执行
	}()
	<-c // 取出channel这个变量的值,阻塞在这里(<-ch 将一直阻塞，直到ch被关闭 或者 ch中可以取出值 为止)
}

// 等同于go4.go
