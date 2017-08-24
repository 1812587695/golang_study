package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			fmt.Printf("ticked at %v", time.Now())
		}
	}()
}
