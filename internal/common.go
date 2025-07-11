package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateSettings() {
	f, err := os.Create("lyra.json")
	if err != nil {
		fmt.Println("Could not create lyra.json")
	}
	defer f.Close()
}

func ProjectRoot() (string, error) {
	dir, _ := os.Getwd()

	for {
		_, err := os.Stat(filepath.Join(dir, "lyra.json"))
		if err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("Project root not found")
}
