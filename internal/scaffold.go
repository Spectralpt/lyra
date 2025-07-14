package internal

import (
	_ "embed"
	"fmt"
	"lyra/globals"
	"lyra/types"
	"os"
	"path/filepath"
	"slices"
)

//go:embed resources/Apache-License.txt
var apacheLicense string

//go:embed resources/GPLv3-License.txt
var gplv3License string

//go:embed resources/MIT-License.txt
var mitLicense string

//go:embed resources/README.md
var readmeTemplate string

//go:embed resources/gitlab-ci.yml
var gitlabCITemplate string

func ScaffoldProject(project types.Project) error {
	var err error

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
		// fmt.Println("creating:", dir)
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

	modules := GetModulesByNames(project.Modules)
	GenerateComposeWithYTT(modules, project.Name)

	// Create LICENSE file based on selected license
	licensePath := filepath.Join(project.Name, "LICENSE")
	switch project.License {
	case "apache":
		err = os.WriteFile(licensePath, []byte(apacheLicense), 0644)
	case "mit":
		err = os.WriteFile(licensePath, []byte(mitLicense), 0644)
	case "gplv3":
		err = os.WriteFile(licensePath, []byte(gplv3License), 0644)
	}
	if err != nil {
		fmt.Printf("err license: %v\n", err)
		return err
	}

	// Create README.md
	readmePath := filepath.Join(project.Name, "README.md")
	err = os.WriteFile(readmePath, []byte(readmeTemplate), 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	// Create GitLab CI config
	gitlabCIPath := filepath.Join(project.Name, ".gitlab-ci.yml")
	err = os.WriteFile(gitlabCIPath, []byte(gitlabCITemplate), 0644)
	if err != nil {
		return err
	}

	return err
}

func GetModulesByNames(names []string) []types.Module {
	var selectedModules []types.Module

	// 1. Always add "fiware-orion" as the first module (base)
	for _, module := range globals.AvailableModules {
		if module.Name == "fiware-orion" {
			selectedModules = append(selectedModules, module)
			break
		}
	}

	for _, module := range globals.AvailableModules {
		if slices.Contains(names, module.Name) {
			selectedModules = append(selectedModules, module)
		}
	}
	return selectedModules
}
