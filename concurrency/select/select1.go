package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	//	close(ch1)
	//	close(ch2)
	//	for {
	select {
	case <-ch1:
		fmt.Println("1111")
	case <-ch2:
		fmt.Println("2222")
	}
	//	}

}
