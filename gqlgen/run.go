package gqlgen

import (
	"github.com/funcgql/cli/go/tools"
	"github.com/pkg/errors"
)

func (a *api) runIn(cmd string, absPath string) error {
	const gqlgenModuleName = "github.com/99designs/gqlgen"

	if output, err := tools.RunIn(gqlgenModuleName, absPath, cmd); err != nil {
		return errors.Wrapf(err, "failed to run gqlgen %s %s", cmd, output.Combined)
	}
	return nil
}
