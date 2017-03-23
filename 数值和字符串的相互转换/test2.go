package main

import (
	"fmt"
	"strconv"
)

func main() {

	d := 1
	d = strconv.Itoa(d) //数字变成字符串,然后赋值给你 整型的num，这时候会报错
	d = "sdfs" + d
	fmt.Println(d)

}
