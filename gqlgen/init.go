package gqlgen

import (
	"github.com/funcgql/cli/go/tools"
	"github.com/pkg/errors"
)

func (a *api) Init(absPath string) error {
	if output, err := tools.RunIn(moduleName, absPath, "init"); err != nil {
		return errors.Wrapf(err, "failed to initialize gqlgen %s", output.Combined)
	}
	return nil
}
