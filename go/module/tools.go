package module

import (
	"path/filepath"
	"regexp"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

func (m module) Tools() ([]Tool, error) {
	const toolsFilename = "tools.go"
	toolsFile := cliio.FileOf(filepath.Join(m.absPath, toolsFilename))
	content, err := toolsFile.StringContent()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s", toolsFile.AbsPath())
	}

	var results []Tool
	importPathRegex := regexp.MustCompile("_ *\"(.+)\"")
	matches := importPathRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			results = append(results, Tool{
				ImportPath:    match[1],
				moduleAbsPath: m.absPath,
			})
		}
	}

	return results, nil
}

func (m module) InstallAllTools(shellAPI shell.API) error {
	tools, err := m.Tools()
	if err != nil {
		return err
	}
	dependencies, err := m.Dependencies(shellAPI)
	if err != nil {
		return err
	}

	for _, tool := range tools {
		// Find tool's version by matching with dependencies.
		dep, hasMatch := dependencies[tool.ImportPath]
		if hasMatch {
			tool.Install(dep.Version, shellAPI)
		} else {
			tool.Install("latest", shellAPI)
		}
	}

	return nil
}
