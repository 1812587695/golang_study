package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	var x float64 = 3.4
	//	v := reflect.ValueOf(x)
	//	fmt.Println(v)
	//	fmt.Println("settability of v:", v.CanSet())
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.注意这里哦！我们把x地址传进去了！
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
