package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取文件权限
	filemode, _ := os.Stat("os1.go")
	fmt.Println(filemode.Mode())      //获取权限 linux 0600
	err := os.Chmod("widuu.go", 0777) //改变的是文件的权限
	if err != nil {
		fmt.Println("修改文件权限失败")
	}
	filemode, _ = os.Stat("os1.go")
	fmt.Println(filemode.Mode()) //获取权限是0777
}
