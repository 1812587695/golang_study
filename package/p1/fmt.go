package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	addr string
}

func main() {
	p := Person{"rain", 23, "qingyangqu"}
	fmt.Println(p)
	a := newCircle()
	b := a.Area()
	fmt.Println(b)
}
