package gqlgen

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/tools"
	lambdatemplate "github.com/funcgql/cli/gqlgen/template/lambda"
	"github.com/pkg/errors"
)

func (a *api) Init(absPath string, moduleName string, functionType functype.FunctionType) error {
	if output, err := tools.RunIn(gqlgenModuleName, absPath, "init"); err != nil {
		return errors.Wrapf(err, "failed to initialize gqlgen %s", output.Combined)
	}

	if err := a.replaceMain(absPath, moduleName, functionType); err != nil {
		return err
	}

	return nil
}

func (a *api) replaceMain(absPath string, moduleName string, functionType functype.FunctionType) error {
	serverFile := cliio.FileOf(filepath.Join(absPath, "server.go"))
	if err := serverFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", serverFile.AbsPath())
	}

	switch functionType {
	case functype.Lambda:
		mainTemplate := lambdatemplate.New(moduleName)
		if err := mainTemplate.Export(absPath); err != nil {
			return errors.Wrap(err, "failed to generate main.go")
		}
	default:
		return errors.Errorf("unknown function type %s to generate main.go", functionType)
	}

	return nil
}
