package main

//#类型"chan<- int"表示一个只接受int的channel，用于数据放置进入
//#类型"<-chan int"表示一个只提取int的channel，用于数据提取离开

import (
	"fmt"
	//	"time"
)

func comein(oneway chan<- string, msg string) {
	fmt.Println(msg, "将要进入单行通道")
	oneway <- msg
}

func goout(oneway <-chan string, mainroad chan<- string) {
	msg := <-oneway
	fmt.Println(msg, "已经离开单行通道")
	fmt.Println(msg, "将要进入主路")
	mainroad <- msg
}

func main() {
	oneway := make(chan string, 1)
	mainroad := make(chan string, 1)
	comein(oneway, "我的车")
	goout(oneway, mainroad)
	//	time.Sleep(2 * time.Second)
	fmt.Println(<-mainroad, "离开主路")
	fmt.Println("是两地分居")
}
