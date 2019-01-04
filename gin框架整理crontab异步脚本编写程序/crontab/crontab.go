package main

import (
	"github.com/robfig/cron"
	"os/exec"
	"fmt"
)

func main() {

	c := cron.New()
	//spec := "0 0 0,1,2 * * ?"  // 每天12点0分，1分，2分执行
	spec := "0 0 0 * * ?" // 每天12点0分执行
	c.AddFunc(spec, func() {
		cmd := exec.Command("./main2")
		cmd.Dir = "/data/gopath/src/gin_sync/"
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	})
	c.Start()

	select{}
}