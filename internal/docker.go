package internal

import (
	_ "embed"
	"fmt"
	"lyra/types"
	"os"
	"os/exec"
	"path/filepath"
)

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

// func GenerateComposeWithYTT(modules []Module, outputPath string) error {
// 	fmt.Println("Modules:", modules)
// 	fmt.Println("Output path:", outputPath)
//
// 	currentDir, _ := os.Getwd()
//
// 	args := []string{}
// 	cmd := exec.Command("ytt", args...)
// 	tempFiles := []string{} // Store temp file paths for cleanup
//
// 	// 1. Create all temp files
// 	for _, overlay := range modules {
// 		file, err := os.CreateTemp(currentDir, "lyra-*.yaml") // Better naming pattern
// 		if err != nil {
// 			return fmt.Errorf("failed to create temp file: %w", err)
// 		}
// 		defer file.Close() // Ensure file handle is closed
//
// 		if _, err := file.WriteString(overlay.Overlay); err != nil {
// 			return fmt.Errorf("failed to write to temp file: %w", err)
// 		}
//
// 		tempFiles = append(tempFiles, file.Name()) // Track for cleanup
// 		args = append(args, "-f", file.Name())
// 	}
//
// 	outputFile, err := os.Create("test.yaml")
// 	if err != nil {
// 		return fmt.Errorf("failed to create output file: %w", err)
// 	}
// 	cmd.Stdout = outputFile
//
// 	fmt.Println("Args:", args)
//
// 	out, err := cmd.Output()
//
// 	fmt.Println("Cmd output:", cmd)
//
// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("command failed: %w", err)
// 	}
//
// 	// 5. Clean up temp files (after command succeeds)
// 	for _, tempFile := range tempFiles {
// 		if err := os.Remove(tempFile); err != nil {
// 			fmt.Printf("Warning: failed to remove temp file %q: %v", tempFile, err)
// 		}
// 	}
// 	defer outputFile.Close()
//
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return fmt.Errorf("ytt failed: %w", err)
// 	}
//
// 	return os.WriteFile(outputPath, out, 0644)
// }

func GenerateComposeWithYTT(modules []types.Module, outputPath string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	var args []string
	tempFiles := []string{} // Store temp file paths for cleanup

	// 1. Create all temp files
	for _, overlay := range modules {
		if overlay.Overlay == "" {
			continue // Skip modules with empty overlays
		}

		file, err := os.CreateTemp(currentDir, "lyra-*.yaml")
		if err != nil {
			return fmt.Errorf("failed to create temp file: %w", err)
		}

		if _, err := file.WriteString(overlay.Overlay); err != nil {
			file.Close()
			return fmt.Errorf("failed to write to temp file: %w", err)
		}
		file.Close() // Close immediately after writing

		tempFiles = append(tempFiles, file.Name())
		args = append(args, "-f", file.Name())
	}

	// 2. Prepare the command
	cmd := exec.Command("ytt", args...)

	// 3. Set up output
	outputFile, err := os.Create(filepath.Join(outputPath, "docker-compose.yaml")) // Use the provided outputPath
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	// 4. Run the command
	if err := cmd.Run(); err != nil {
		// Clean up temp files even if command fails
		for _, tempFile := range tempFiles {
			if err := os.Remove(tempFile); err != nil {
				fmt.Printf("Warning: failed to remove temp file %q: %v\n", tempFile, err)
			}
		}
		return fmt.Errorf("ytt command failed: %w", err)
	}

	// 5. Clean up temp files
	for _, tempFile := range tempFiles {
		if err := os.Remove(tempFile); err != nil {
			return err
		}
	}

	return nil
}
