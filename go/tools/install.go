package tools

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/funcgql/cli/cliio"
	"github.com/pkg/errors"
)

func (a *api) InstallAllIn(dir string) error {
	tools, err := goToolsIn(dir)
	if err != nil {
		return err
	}
	for _, tool := range tools {
		for _, toolName := range tool.toolNames {
			fmt.Println("🐭 Installing", toolName, "in", tool.modulePath)
			if output, err := a.shellAPI.ExecuteIn(tool.modulePath, "go", "install", toolName); err != nil {
				return errors.Wrapf(err, "failed to install %s in %s %s", toolName, tool.modulePath, output.Combined)
			}
		}
	}

	return nil
}

type goTools struct {
	modulePath string
	toolNames  []string
}

func goToolsIn(dir string) ([]goTools, error) {
	toolsFilePaths, err := goToolsFilePathsIn(dir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all tools.go files")
	}

	var results []goTools
	for _, toolsFilePath := range toolsFilePaths {
		toolNames, err := goToolsModuleNamesIn(toolsFilePath)
		if err != nil {
			return nil, err
		}
		results = append(results, goTools{
			modulePath: filepath.Dir(toolsFilePath),
			toolNames:  toolNames,
		})
	}

	return results, nil
}

func goToolsFilePathsIn(dir string) ([]string, error) {
	var results []string
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d != nil && d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "tools.go") {
			results = append(results, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return results, nil
}

func goToolsModuleNamesIn(path string) ([]string, error) {
	content, err := cliio.FileOf(path).StringContent()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read go tools file at %s", path)
	}

	var moduleNames []string
	moduleNameRegex := regexp.MustCompile("_ *\"(.+)\"")
	matches := moduleNameRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			moduleNames = append(moduleNames, match[1])
		}
	}

	return moduleNames, nil
}
