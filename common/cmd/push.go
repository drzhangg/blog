package cmd

import (
	"log"
	"os/exec"
)

func PushGithub(commit string) error {
	cmd := exec.Command("/bin/sh", "git.sh", commit)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("执行完毕")
	return nil
}
