package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())      //显示当前的目录 D:\test <nil>
	fmt.Println(os.Chdir("D:/")) //返回<nil>正确切换目录了
	fmt.Println(os.Getwd())      //切换后的目录D:\ <nil>
}
