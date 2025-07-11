package internal

import (
	"fmt"
	"lyra/internal/types"
	"os"
	"path/filepath"
)

func ScaffoldProject(project types.Project) error {
	var err error
	fmt.Println(project)

	//testing
	if _, err := os.Stat(project.Name); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", project.Name)
	}
	if err := os.Mkdir(project.Name, 0755); err != nil {
		return err
	} //testing
	dirs := [...]string{
		"config",
		"services",
	}

	for _, dir := range dirs {
		fmt.Println("creating:", dir)
		err := os.Mkdir(filepath.Join(project.Name, dir), 0755)
		if err != nil {
			return err
		}
	}

	if project.Git {
		path, _ := os.Getwd()
		err = GitInit(filepath.Join(path, project.Name))
		if err != nil {
			return err
		}
	}

	CreateCompose(project.Name)

	return err
}
