package main

/**
* 这是第9个go concurrency 实例
 */

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// 加不加这句话执行完时间都是一致的
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
	fmt.Println(time.Second)
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
