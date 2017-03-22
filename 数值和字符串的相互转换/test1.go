package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 10; i++ {

		d := strconv.Itoa(i) //数字变成字符串
		d = "sdfs" + d
		fmt.Println(d)
	}

}
