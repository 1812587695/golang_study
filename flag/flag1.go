package main

//标准库—命令行参数解析flag

import (
	"flag"
	"fmt"
)

// file是指定参数：如（go run flag1.go -file "../music/sun.mp3"）
var music_file *string = flag.String("file", "默认返回值", "报错返回信息")

func main() {
	flag.Parse()
	fmt.Println(*music_file)
}

/**
用法
go run flag1.go -file "../music/sun.mp3"
go run flag1.go，musicfile是默认值

*/
