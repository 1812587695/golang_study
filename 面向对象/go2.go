package main

//继承特性
import "fmt"

type parent struct {
	val int
}

type child struct {
	parent
	num int
}

func main() {
	var c child

	c = child{parent{1}, 2}
	fmt.Println(c.num)
	fmt.Println(c.val)
}
