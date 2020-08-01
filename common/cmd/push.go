package cmd

import "os/exec"

func PushGithub(commit string) error {
	cmd := exec.Command("/bin/sh", "git.sh", commit)
	err := cmd.Run()
	if err != nil {

		return err
	}
	return nil
}
