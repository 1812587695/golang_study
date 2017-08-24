package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	locker := sync.Mutex{}
	cond := sync.NewCond(&locker)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cond.Signal()                  //当键盘输入enter后,发出通知信号
		fmt.Println("signal...")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("sleep end...")

	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock() //首先进行锁定,与之关联的条件变量的锁定
		fmt.Println("wait before...")
		//等待Cond消息通知
		cond.Wait()
		fmt.Println("wait end...")
		cond.L.Unlock()
	}()

	wg.Wait()
	fmt.Println("exit...")
}
