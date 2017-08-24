package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var num int

	for i := 1; i <= 2; i++ {
		go func(id int) {
			fmt.Println("Enter Thread ID:", id)
			c.L.Lock()
			for num != 1 {
				fmt.Println("Enter loop: Thread ID:", id)
				c.Wait()
				fmt.Println("Exit loop: Thread ID:", id)
			}
			num++
			c.L.Unlock()
			fmt.Println("Exit Thread ID:", id)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Sleep 1 second")

	num++
	c.Broadcast()
	time.Sleep(time.Second)
	fmt.Println("Program exit")
}
