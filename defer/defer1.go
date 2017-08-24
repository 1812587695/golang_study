package main

import (
	"fmt"
)

//defer是在return之前执行的

func main() {
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f3())
}
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
