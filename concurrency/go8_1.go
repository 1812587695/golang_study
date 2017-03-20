package main

/**
* 比较第8个go concurrency 实例运行时间
 */

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		Go(i)
	}

	fmt.Println(time.Now())
}

func Go(index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)
}
