package module

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type Tool struct {
	ImportPath    string
	moduleAbsPath string
}

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

func (t Tool) Install(version string, shellAPI shell.API) error {
	versionedImportPath := fmt.Sprintf("%s@%s", t.ImportPath, version)
	if output, err := shellAPI.ExecuteIn(t.moduleAbsPath, "go", "install", versionedImportPath); err != nil {
		return errors.Wrapf(err, "failed to install %s in %s %s", versionedImportPath, t.moduleAbsPath, output.Combined)
	}

	return nil
}
