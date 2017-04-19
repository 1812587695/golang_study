package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var M map[string]string
	File, _ := os.Open("test.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}

