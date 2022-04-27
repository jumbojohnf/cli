package module

import (
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

func (m module) Tidy(shellAPI shell.API) error {
	if output, err := shellAPI.ExecuteWithIOIn(m.absPath, "go", "mod", "tidy"); err != nil {
		return errors.Wrap(err, output.Combined)
	}
	return nil
}
