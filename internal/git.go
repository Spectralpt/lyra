package internal

import (
	"errors"
	"os/exec"
)

func isGitInstalled() bool {
	_, err := exec.LookPath("git")
	if err != nil {
		return false
	}
	return true
}

func GitInit(dir string) error {
	if !isGitInstalled() {
		return errors.New("git is not installed")
	}
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		println("from gitinit:", err.Error())
		return err
	}
	return nil
}

// Todo
func AddRemote(remote string) error {
	return nil
}
