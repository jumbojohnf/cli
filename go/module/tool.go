package module

import (
	"fmt"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type Tool struct {
	ImportPath    string
	moduleAbsPath string
}

func (t Tool) Install(version string, shellAPI shell.API) error {
	versionedImportPath := fmt.Sprintf("%s@%s", t.ImportPath, version)
	if output, err := shellAPI.ExecuteIn(t.moduleAbsPath, "go", "install", versionedImportPath); err != nil {
		return errors.Wrapf(err, "failed to install %s in %s %s", versionedImportPath, t.moduleAbsPath, output.Combined)
	}

	return nil
}
