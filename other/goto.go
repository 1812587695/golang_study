package main

import (
	"fmt"
)

var t int = 2

func main() {
	if t > 10 {
		goto A
	} else {
		goto B
	}
A:
	{
		fmt.Println("A")
		goto C
	}

B:
	{
		fmt.Println("B")
		goto C
	}
C:
}
