package main

/**
* 这是第8个go concurrency 实例
 */

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10) // 这里循环10次，我设置10个缓存
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	// 循环读取10次缓存channel变量c的值，阻塞在这里，如果10次读完程序才执行完毕
	for i := 0; i < 10; i++ {
		<-c
	}
	fmt.Println(time.Second)
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true

}
