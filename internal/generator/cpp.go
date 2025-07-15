package generator

import (
	_ "embed"
	"os"
	"path/filepath"
)

// Embedding all the necessary resources for the C++ project
//
//go:embed resources/cpp/test_main.cpp
var cppTestMainContent string

//go:embed resources/cpp/Makefile
var cppMakefileContent string

//go:embed resources/cpp/CMakeLists.txt
var cppCMakeContent string

//go:embed resources/cpp/main.cpp
var cppMainContent string

//go:embed resources/cpp/utils.cpp
var cppUtilsContent string

//go:embed resources/cpp/utils.hpp
var cppUtilsHeaderContent string

//go:embed resources/cpp/.gitignore
var cppGitignoreContent string

//go:embed resources/cpp/CPP_README.md
var cppReadmeContent string

type CppGenerator struct{}

// GenerateTests creates the test files for the C++ project
func (cpp *CppGenerator) GenerateTests() error {
	// Create tests directory
	if err := os.MkdirAll("tests", 0755); err != nil {
		return err
	}

	// Create test_main.cpp file
	err := os.WriteFile(filepath.Join("tests", "test_main.cpp"), []byte(cppTestMainContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

// GenerateBuildSystemConfig sets up the build system configuration files
func (cpp *CppGenerator) GenerateBuildSystemConfig() error {
	// Create build directory
	if err := os.MkdirAll("build", 0755); err != nil {
		return err
	}

	// Create Makefile
	err := os.WriteFile("Makefile", []byte(cppMakefileContent), 0644)
	if err != nil {
		return err
	}

	// Create CMakeLists.txt as an alternative build system
	err = os.WriteFile("CMakeLists.txt", []byte(cppCMakeContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

// GenerateAditionalScafolding creates the additional scaffolding for the C++ project
func (cpp *CppGenerator) GenerateAditionalScafolding() error {
	// Create services/src directory
	if err := os.MkdirAll(filepath.Join("services", "src"), 0755); err != nil {
		return err
	}

	// Create main.cpp in services/src
	err := os.WriteFile(filepath.Join("services", "src", "main.cpp"), []byte(cppMainContent), 0644)
	if err != nil {
		return err
	}

	// Create utils.cpp in services/src
	err = os.WriteFile(filepath.Join("services", "src", "utils.cpp"), []byte(cppUtilsContent), 0644)
	if err != nil {
		return err
	}

	// Create utils.hpp header file
	err = os.WriteFile(filepath.Join("services", "src", "utils.hpp"), []byte(cppUtilsHeaderContent), 0644)
	if err != nil {
		return err
	}

	// Create a .gitignore for C++ projects
	err = os.WriteFile(".gitignore", []byte(cppGitignoreContent), 0644)
	if err != nil {
		return err
	}

	// Create a README for the C++ project structure
	err = os.WriteFile("CPP_README.md", []byte(cppReadmeContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
