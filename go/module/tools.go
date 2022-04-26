package module

import (
	"path/filepath"
	"regexp"

	"github.com/funcgql/cli/cliio"
	"github.com/pkg/errors"
)

type Tool struct {
	ImportPath string
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
			results = append(results, Tool{ImportPath: match[1]})
		}
	}

	return results, nil
}
