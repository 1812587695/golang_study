package main

import (
	"fmt"
	"sync"
)

var count = 0
var condition = 0

func main() {
	lock := new(sync.Mutex)
	cond := sync.NewCond(lock)

	go func() {
		for {
			lock.Lock()
			for condition == 0 {
				cond.Wait()
			}
			fmt.Printf("Consumed %d\n", count)
			condition = 0
			cond.Signal()
			lock.Unlock()
		}
	}()
	for {
		lock.Lock()
		for condition == 1 {
			cond.Wait()
		}
		fmt.Printf("Produced %d\n", count)
		count++
		condition = 1
		cond.Signal()
		lock.Unlock()
	}
}
