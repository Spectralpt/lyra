package generator

import "fmt"

type Generator interface {
	GenerateTests() error
	GenerateBuildSystemConfig() error
	GenerateAditionalScafolding() error
}

func NewGenerator(language string) (Generator, error) {
	switch language {
	case "c":
		return &CGenerator{}, nil
	case "cpp":
		return &CppGenerator{}, nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}
