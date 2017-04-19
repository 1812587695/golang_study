package main

import "fmt"

import "io/ioutil"

func main() {

	dir_list, e := ioutil.ReadDir("C:/Users/Administrator/go/src/golang_study/io")

	if e != nil {

		fmt.Println("read dir error")

		return

	}

	// 读取目录下所有的文件和目录
	for i, v := range dir_list {

		fmt.Println(i, "=", v.Name())

		fmt.Println(v.Name(), "的权限是:", v.Mode())

		fmt.Println(v.Name(), "文件大小:", v.Size())

		fmt.Println(v.Name(), "创建时间", v.ModTime())

		fmt.Println(v.Name(), "系统信息", v.Sys())

		if v.IsDir() == true {

			fmt.Println(v.Name(), "是目录")

		}

	}

}
