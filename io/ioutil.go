package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 读取目录
	rd, err := ioutil.ReadDir("D:/")
	fmt.Println(err)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", fi.Name())

		} else {
			fmt.Println(fi.Name())
		}
	}
}
