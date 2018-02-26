package main

import (
	"fmt"
)

type dog struct{

}

func main() {
	var d []dog
	fmt.Println(d)
	d1 := new(dog)
	fmt.Println(d1)
}
