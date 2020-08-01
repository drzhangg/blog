package main

import "os/exec"

func main() {
	cmd := exec.Command("/bin/sh", "git.sh", "git")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
