package main

/**
* 这是第二个go concurrency 实例
  解决上go1.go的疑问
*/

import (
	"fmt"
	"time"
)

func main() {
	go Hello()
	// 当时间在这里停顿了2秒钟，Hello函数就打印输出了内容，就简单的实现了阻塞等待，
	// 为什么我知道hello函数在2秒钟以内可以运行输出完成呢？如果是其他函数，我是不是修改10秒，或者更长时间呢？显然不行.
	time.Sleep(2 * time.Second)
}

func Hello() {
	fmt.Println("hello wrod!!!")
}

// go3.go解决疑问
