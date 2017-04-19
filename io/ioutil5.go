package main

import (
	"fmt"

	"io/ioutil"
)

func main() {

	// 创建文件
	dir, err := ioutil.TempDir("C:/Users/Administrator/go/src/golang_study/io", "tmp")

	if err != nil {

		fmt.Println("常见临时目录失败")

		return

	}

	fmt.Println(dir) //返回的是D:\test\tmp846626247 就是前边的prefix+随机数

}
