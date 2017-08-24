package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type T struct {
	l *sync.Mutex // 锁
	c *sync.Cond  //条件变量
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	var t *T = new(T)
	t.l = new(sync.Mutex)
	// 使用条件变量前，必须将其与一个锁绑定
	t.c = sync.NewCond(t.l)

	//启动10条协程, 当协程执行conditions()发现满足条件时，打印出processed，代表该协程处理结束
	//当不满足条件时，通过条件变量进入睡眠状态，每次从睡眠状态醒来并且获取锁之后，打印出一条wait。
	for i := 0; i < 10; i++ {
		go func() {
			t.l.Lock()
			defer t.l.Unlock()
			//不满足条件时通过条件变量进入沉入
			//t.c.Wait() 首先会释放与该条件变量绑定的锁，然后在进入睡眠状态
			for !conditions() {
				t.c.Wait()
				fmt.Println("wait")
			}
			fmt.Println("processed")

		}()
	}

	// 启动一条协程 每隔两秒发送一次通知
	go func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("\nBroadcast")
			t.c.Broadcast()
			//			fmt.Println("\nSignal")
			//			t.c.Signal()
		}

	}()
	wg.Wait()

}

//生成随机数， 当生成的数为0时，则为满足条件返回true
func conditions() bool {
	i := rand.Intn(10)
	fmt.Println(i)
	if i == 0 {
		return true
	} else {
		return false
	}
}
