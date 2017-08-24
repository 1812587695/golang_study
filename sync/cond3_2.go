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

	var condition bool = false

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cond.L.Lock()
		cond.Signal()    //当键盘输入enter后,发出通知信号
		condition = true //把条件变量设为true,表示发送过信号
		fmt.Println("signal...")
		cond.L.Unlock()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("sleep end...")

	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock() //首先进行锁定,与之关联的条件变量的锁定
		fmt.Println("wait before...")
		//等待Cond消息通知
		for !condition {
			//当条件为真时,不会发生wait
			fmt.Println("wait...")
			cond.Wait()
		}
		fmt.Println("wait end...")
		cond.L.Unlock()
	}()

	wg.Wait()
	fmt.Println("exit...")
}
