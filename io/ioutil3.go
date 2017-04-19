package main

import (
	"fmt"

	"io/ioutil"

	"reflect"

	"strings"
)

func main() {

	reader := strings.NewReader("hello word widuu") //返回*strings.Reader

	fmt.Println(reflect.TypeOf(reader))

	data, _ := ioutil.ReadAll(reader)

	fmt.Println(string(data))

}
