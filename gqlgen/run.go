package gqlgen

import (
	"github.com/funcgql/cli/go/module"
	"github.com/pkg/errors"
)

func (a *api) run(cmd string, targetModule module.Module) error {
	const gqlgenImportPath = "github.com/99designs/gqlgen"
	gqlgenTool := targetModule.ToolOf(gqlgenImportPath)

	if output, err := gqlgenTool.Run(a.shellAPI, cmd); err != nil {
		return errors.Wrapf(err, "failed to run gqlgen %s %s", cmd, output.Combined)
	}
	return nil
}
