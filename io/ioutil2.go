package main

import (
	"fmt"

	"io/ioutil"

	"os"
	"reflect"
)

func main() {

	// 读取文件的内容
	data, err := ioutil.ReadFile("C:/Users/Administrator/go/src/golang_study/io/ioutil.go")

	if err != nil {

		fmt.Println("read error")

		os.Exit(1)

	}

	fmt.Println(string(data))

	// 反射类型
	fmt.Println("type:", reflect.TypeOf(data))

}
