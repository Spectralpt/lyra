package internal

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed resources/Apache-License.txt
var apacheLicense string

//go:embed resources/GPLv3-License.txt
var gplv3License string

//go:embed resources/MIT-License.txt
var mitLicense string

//go:embed resources/Makefile
var makefileTemplate string

//go:embed resources/README.md
var readmeTemplate string

//go:embed resources/gitlab-ci.yml
var gitlabCITemplate string

type ProjectConfig struct {
	ProjectName string
	LicenseType string // "apache", "mit", or "gplv3"
	//Language    string // "c" or "cpp"
}

func CProjectScaffold(config ProjectConfig) error {
	baseDir, err := ProjectRoot()
	if err != nil {
		return err
	}

	// Create directories
	dirs := [4]string{"tests", "lib", "docs", "src"}

	for _, dir := range dirs {
		err = os.Mkdir(filepath.Join(baseDir, dir), 0755)
		if err != nil {
			return err
		}
	}

	// Create LICENSE file based on selected license
	switch config.LicenseType {
	case "apache":
		err = os.WriteFile(baseDir, []byte(apacheLicense), 0644)
	case "mit":
		err = os.WriteFile(baseDir, []byte(mitLicense), 0644)
	case "gplv3":
		err = os.WriteFile(baseDir, []byte(gplv3License), 0644)
	default:
		err = os.WriteFile(baseDir, []byte(apacheLicense), 0644) // default to Apache
	}
	if err != nil {
		return err
	}

	// Create README.md
	readmePath := filepath.Join(baseDir, "README.md")
	err = os.WriteFile(readmePath, []byte(readmeTemplate), 0644)
	if err != nil {
		return err
	}

	// Create Makefile
	makefilePath := filepath.Join(baseDir, "Makefile")
	err = os.WriteFile(makefilePath, []byte(makefileTemplate), 0644)
	if err != nil {
		return err
	}

	// Create GitLab CI config
	gitlabCIPath := filepath.Join(baseDir, ".gitlab-ci.yml")
	err = os.WriteFile(gitlabCIPath, []byte(gitlabCITemplate), 0644)
	if err != nil {
		return err
	}

	// Create LICENSE file based on selected license
	switch config.LicenseType {
	case "apache":
		err = os.WriteFile(baseDir, []byte(apacheLicense), 0644)
	case "mit":
		err = os.WriteFile(baseDir, []byte(mitLicense), 0644)
	case "gplv3":
		err = os.WriteFile(baseDir, []byte(gplv3License), 0644)
	default:
		err = os.WriteFile(baseDir, []byte(apacheLicense), 0644) // default to Apache
	}
	if err != nil {
		return err
	}

	return nil
}
