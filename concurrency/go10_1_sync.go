package main

import (
	"log"

	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // 10行这里取地址符号，如果这里不取，在17行也得取地址符号

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			log.Printf("i:%d", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	log.Println("exit")
}
