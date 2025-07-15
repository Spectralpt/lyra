package generator

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed resources/c/test_main.c
var cTestMainContent string

//go:embed resources/c/Makefile
var cMakefileContent string

//go:embed resources/c/main.c
var cMainContent string

//go:embed resources/c/utils.c
var cUtilsContent string

//go:embed resources/c/utils.h
var cUtilsHeaderContent string

//go:embed resources/c/.gitignore
var cGitignoreContent string

//go:embed resources/c/C_README.md
var cReadmeContent string

type CGenerator struct{}

func (c *CGenerator) GenerateTests() error {
	// Create tests directory
	if err := os.MkdirAll("tests", 0755); err != nil {
		return err
	}

	err := os.WriteFile(filepath.Join("tests", "test_main.c"), []byte(cTestMainContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *CGenerator) GenerateBuildSystemConfig() error {
	// Create build directory
	if err := os.MkdirAll("build", 0755); err != nil {
		return err
	}

	err := os.WriteFile("Makefile", []byte(cMakefileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *CGenerator) GenerateAditionalScafolding() error {
	// Create services/src directory
	if err := os.MkdirAll(filepath.Join("services", "src"), 0755); err != nil {
		return err
	}

	// Create main.c in services/src
	err := os.WriteFile(filepath.Join("services", "src", "main.c"), []byte(cMainContent), 0644)
	if err != nil {
		return err
	}

	// Create utils.c in services/src
	err = os.WriteFile(filepath.Join("services", "src", "utils.c"), []byte(cUtilsContent), 0644)
	if err != nil {
		return err
	}

	// Create utils.h header file
	err = os.WriteFile(filepath.Join("services", "src", "utils.h"), []byte(cUtilsHeaderContent), 0644)
	if err != nil {
		return err
	}

	// Create a .gitignore for C projects
	err = os.WriteFile(".gitignore", []byte(cGitignoreContent), 0644)
	if err != nil {
		return err
	}

	// Create a README for the C project structure
	err = os.WriteFile("C_README.md", []byte(cReadmeContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
