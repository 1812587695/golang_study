package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3e9) // sleep one second
		timeout <- true
	}()

	ch := make(chan int)

	go func() {
		ch <- 222
	}()

	select {
	case a := <-ch:
		fmt.Println(a)
	case <-timeout:
		fmt.Println("timeout!")
	}

}
