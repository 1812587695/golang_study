package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc(i int) {
	//time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!", i)
}

func main() {
	defer ants.Release()

	runTimes := 10000000

	// Use the common pool
	var wg sync.WaitGroup
	/*	for i := 0; i < runTimes; i++ {
			wg.Add(1)
			fmt.Println(i)
			ants.Submit(func() {demoFunc(i);wg.Done()})
		}
		wg.Wait()
		fmt.Printf("running goroutines: %d\n", ants.Running())
		fmt.Printf("finish all tasks.\n")*/

	// Use the pool with a function,
	// set 10 to the size of goroutine pool and 1 second for expired duration

	p, _ := ants.NewPoolWithFunc(10000000, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Serve(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}