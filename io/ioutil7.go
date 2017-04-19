package main

import (
	"fmt"

	"io/ioutil"
)

func main() {
	// 创建文件
	file, error := ioutil.TempFile("C:/Users/Administrator/go/src/golang_study/io", "tmp.txt")

	defer file.Close()

	if error != nil {

		fmt.Println("创建文件失败")

		return

	}
	filedata_w, _ := ioutil.ReadFile("C:/Users/Administrator/go/src/golang_study/io/ioutil.go")
	// 写入文件内容
	file.WriteString(string(filedata_w)) //利用file指针的WriteString()详情见os.WriteString()

	// 读取文件内容
	filedata, _ := ioutil.ReadFile(file.Name())

	fmt.Println(string(filedata))

}
