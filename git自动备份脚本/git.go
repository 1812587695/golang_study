package main

import (
	"os/exec"
	"time"
)

func exec_shell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)

	cmd.Run()
}

func main() {
	for {
		exec_shell("git pull &>>gitpull.log")
		time.Sleep(2 * time.Second)
	}
}
