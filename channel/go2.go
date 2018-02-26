package main

import "fmt"

type act interface {
	write()
}

type xiaoming struct {
}

type xiaofang struct {
}

func (xm *xiaoming) write() {
	fmt.Println("xiaoming write")
}

func (xf *xiaofang) write() {
	fmt.Println("xiaofang write")
}

func main() {
	var x act
	//	var m xiaoming
	m := new(xiaoming)
	f := xiaofang{}

	x = m
	x.write()

	x = &f
	x.write()
}
