package module

import (
	"strings"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type Dependency struct {
	ImportPath string
	Version    string
}

func (m module) Dependencies(shellAPI shell.API) (map[string]Dependency, error) {
	output, err := shellAPI.ExecuteIn(m.absPath, "go", "list", "-m", "all")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list dependencies for %s", m.name)
	}
	data := strings.Trim(output.Stdout, " \n")

	results := map[string]Dependency{}
	for _, line := range strings.Split(data, "\n") {
		depData := strings.Split(line, " ")
		if len(depData) == 0 {
			continue
		} else if len(depData) == 1 {
			importPath := depData[0]
			results[importPath] = Dependency{
				ImportPath: importPath,
			}
		} else {
			importPath := depData[0]
			results[importPath] = Dependency{
				ImportPath: importPath,
				Version:    depData[1],
			}
		}
	}
	return results, nil
}
