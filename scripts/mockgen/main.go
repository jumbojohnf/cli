package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("missing mockgen config file path")
	}

	shellAPI := shell.NewAPI()

	configPath := os.Args[1]
	configData, err := loadConfigFrom(configPath)
	if err != nil {
		return err
	}

	for _, target := range mockTargetPackagesFrom(configData) {
		if err := target.generate(shellAPI); err != nil {
			return err
		}
	}

	return nil
}

func loadConfigFrom(path string) (map[string][]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read content of config file at %s", path)
	}

	var result map[string][]string
	if err := yaml.Unmarshal(bytes, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config file content")
	}

	return result, nil
}

const (
	mockGen          = "github.com/golang/mock/mockgen"
	moduleImportPath = "github.com/funcgql/cli"
	mocksPackageName = "mocks"
)

type mockTargetPackage struct {
	packageRelPath    string
	packageImportPath string
	interfaceNames    []string
}

func mockTargetPackagesFrom(data map[string][]string) []mockTargetPackage {
	var results []mockTargetPackage
	for packageRelPath, interfaceNames := range data {
		results = append(results, mockTargetPackage{
			packageRelPath:    packageRelPath,
			packageImportPath: filepath.Join(moduleImportPath, packageRelPath),
			interfaceNames:    interfaceNames,
		})
	}

	return results
}

func (p mockTargetPackage) generate(shellAPI shell.API) error {
	for _, interafaceName := range p.interfaceNames {
		fmt.Println("🥸  Generating mocks for", p.packageImportPath, interafaceName)

		if _, err := shellAPI.Execute(
			"go", "run", mockGen,
			fmt.Sprintf("-destination=%s/%s/mock_%s.go", p.packageRelPath, mocksPackageName, strings.ToLower(interafaceName)),
			fmt.Sprintf("-package=%s", mocksPackageName),
			p.packageImportPath,
			interafaceName,
		); err != nil {
			return err
		}
	}

	return nil
}
