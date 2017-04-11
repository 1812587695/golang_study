package main

import (
	"log"

	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			log.Printf("i:%d", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	log.Println("exit")
}
