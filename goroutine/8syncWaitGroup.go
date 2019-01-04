package main

import (
	"fmt"
	"sync"
)

func main() {


	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Go(&wg, i)
	}

	// 阻塞在这里
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
