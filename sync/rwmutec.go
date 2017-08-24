package main

//golang读写锁RWMutex
//基本遵循两大原则：

//1、可以随便读，多个goroutine同时读
//2、写的时候，啥也不能干。不能读也不能写
//RWMutex提供了四个方法：

//func (*RWMutex) Lock // 写锁定
//func (*RWMutex) Unlock // 写解锁

//func (*RWMutex) RLock // 读锁定
//func (*RWMutex) RUnlock // 读解锁

import (
	"sync"

	"time"
)

var m *sync.RWMutex

func main() {

	m = new(sync.RWMutex)

	//可以多个同时读

	go read(1)

	go read(2)

	time.Sleep(2 * time.Second)

}

func read(i int) {

	println(i, "开始读")

	m.RLock()

	println(i, "正在读。。。")

	time.Sleep(1 * time.Second)

	m.RUnlock()

	println(i, "读完毕。")

}
