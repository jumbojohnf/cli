package module

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/go/module/template/gomod"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

// Initial tools are parsed by matching the tools.go file against the go.mod file. This is necessary since the
// module's dependencies cannot be obtained by running go list -m all before the module is tidied up. And the
// module cannot be tidied up until the initial code and tools are setup.
func (m module) InstallInitialTools(shellAPI shell.API) error {
	tools, err := m.Tools()
	if err != nil {
		return err
	}
	dependencies, err := m.dependenciesFromModFile()
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

func (m module) dependenciesFromModFile() (map[string]Dependency, error) {
	modFile := cliio.FileOf(filepath.Join(m.absPath, gomod.Filename))
	modFileContent, err := modFile.StringContent()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read go.mod file at %s", modFile.AbsPath())
	}

	results := map[string]Dependency{}
	depRegex := regexp.MustCompile("(.+) +(v[^ \n]+)")
	matches := depRegex.FindAllStringSubmatch(modFileContent, -1)
	for _, match := range matches {
		if len(match) > 2 {
			importPath := strings.TrimSpace(match[1])
			version := strings.TrimSpace(match[2])
			results[importPath] = Dependency{
				ImportPath: importPath,
				Version:    version,
			}
		}
	}

	return results, nil
}
