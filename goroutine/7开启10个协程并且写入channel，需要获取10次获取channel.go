package main

/**
* 开启10个协程并且写入channel，需要获取10次获取channel
 */

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	// 开启10个协程
	for i := 0; i < 10; i++ {
		go run(c, i)
	}

	// 因为开启了10次协程，需要十次获取channel，如果少一次就会少执行一次goruntine
	<-c
	<-c
	<-c
	<-c
	<-c
	<-c
	<-c
	<-c
	<-c
	<-c
}

func run(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}

