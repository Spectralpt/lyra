package utils

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

func GitInit(name string) error {
	if !isGitInstalled() {
		return errors.New("git is not installed")
	}
	cmd := exec.Command("git", "init")
	err := cmd.Err
	if err != nil {
		return err
	}
	return nil
}

// Todo
func AddRemote(remote string) error {
	return nil
}
