package main

//封装特性
import "fmt"

type data struct {
	val int
}

func (p_data *data) set(num int) {
	p_data.val = num
}

func (p_data *data) show() {
	fmt.Println(p_data.val)
}

func main() {
	p_data := &data{4}
	p_data.set(5)
	p_data.show()
}
