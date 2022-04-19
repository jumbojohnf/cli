package gqlgen

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/go/tools"
	"github.com/pkg/errors"
)

func (a *api) Init(absPath string) error {
	if output, err := tools.RunIn(moduleName, absPath, "init"); err != nil {
		return errors.Wrapf(err, "failed to initialize gqlgen %s", output.Combined)
	}

	return nil
}

func (a *api) replaceMain(absPath string) error {
	cliio.FileOf(filepath.Join(absPath, "server.go"))
	return nil
}
