package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ComposeModule interface {
	Name() string
	ComposeFragment() map[string]any
	Extend(others map[string]any) error
}

func IsDockerInstalled() bool {
	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Docker is installed")
	}
	return true
}

func CreateComposeAfter() error {
	dir, err := ProjectRoot()
	if err != nil {
		fmt.Println("Couldnt find a lyra project")
		return err
	}
	path := filepath.Join(dir, "docker-compose.yml")
	fmt.Printf("Path:%v\n", path)
	_, err = os.Stat(path)
	if os.IsExist(err) {
		fmt.Println("Docker Compose already exists")
		return err
	} else {
		f, err := os.Create(path)
		fmt.Println(err)
		if err != nil {
			fmt.Println("Could not create docker-compose.yml")
			return err
		}
		defer f.Close()
	}
	return err
}

func CreateCompose(path string) error {
	f, err := os.Create(filepath.Join(path, "docker-compose.yaml"))
	if err != nil {
		fmt.Println("Could not create docker-compose.yml")
		return err
	}
	defer f.Close()
	return err
}
