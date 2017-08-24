package main

import (
	"fmt"
	"runtime"
)

func main() {

	name := "啊啊"
	go func() {

		fmt.Println(":", name)
	}()
	runtime.Gosched()
	name = "版本"
}
