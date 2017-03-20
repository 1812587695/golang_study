package main

/**
* 这是第7个go concurrency 实例
 */

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<-c
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}

}

// 协程的调度是随机的.当执行第index=9的时候，程序就退出，可能没有全部执行完毕
// 这个时候如果需要协成全部循环10次请看go8.go分解
