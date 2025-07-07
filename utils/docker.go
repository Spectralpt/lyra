package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	// "github.com/compose-spec/compose-go/v2/cli"
)

func isDockerInstalled() bool {
	_, err := exec.LookPath("docker")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Docker is installed")
	}
	return true
}

func createCompose() error {
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

// func main() {
// 	dir, _ := projectRoot()
// 	composeFilePath := filepath.Join(dir, "docker-compose.yml")
// 	ctx := context.Background()
// 	options, err := cli.NewProjectOptions([]string{composeFilePath})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	project, err := options.LoadProject(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	projectYaml, err := project.MarshalYAML()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(projectYaml))
// }
