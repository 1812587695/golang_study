package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	a = "123456"
	b, error := strconv.Atoi(a)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	b = b + 1
	fmt.Println(b)
	var c int = 1234
	d := strconv.Itoa(c) //数字变成字符串
	d = d + "sdfs"
	fmt.Println(d)
	var e interface{}
	e = 10
	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		break
	case string:
		fmt.Println("字符串", v)
		break

	}
}
