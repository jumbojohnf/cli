package gqlgen

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
	lambdatemplate "github.com/funcgql/cli/gqlgen/template/lambda"
	localtemplate "github.com/funcgql/cli/gqlgen/template/local"
	schematemplate "github.com/funcgql/cli/gqlgen/template/schema"
	"github.com/pkg/errors"
)

func (a *api) Init(targetModule module.Module, functionTypes []functype.FunctionType) error {
	if err := a.run("init", targetModule); err != nil {
		return err
	}

	if err := a.replaceMain(targetModule, functionTypes); err != nil {
		return err
	}
	if err := a.replaceSchema(targetModule); err != nil {
		return err
	}
	if err := a.replaceResolvers(targetModule); err != nil {
		return err
	}

	return nil
}

func (a *api) replaceMain(targetModule module.Module, functionTypes []functype.FunctionType) error {
	serverFile := cliio.FileOf(filepath.Join(targetModule.AbsPath(), "server.go"))
	if err := serverFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", serverFile.AbsPath())
	}

	for _, functionType := range functionTypes {
		switch functionType {
		case functype.Lambda:
			mainTemplate := lambdatemplate.New(targetModule.Name())
			if err := mainTemplate.Export(targetModule.AbsPath()); err != nil {
				return errors.Wrap(err, "failed to generate lambda server main.go")
			}
		default:
			return errors.Errorf("unknown function type %s to generate main.go", functionType)
		}
	}

	// Regardless of function type, always generate local main.go.
	localMainTemplate := localtemplate.New(targetModule.Name())
	if err := localMainTemplate.Export(targetModule.AbsPath()); err != nil {
		return errors.Wrap(err, "failed to generate local server main.go")
	}

	return nil
}

func (a *api) replaceSchema(targetModule module.Module) error {
	schemaFile := cliio.FileOf(filepath.Join(targetModule.AbsPath(), schematemplate.DirName, schematemplate.Filename))
	if err := schemaFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", schemaFile.AbsPath())
	}

	schemaTemplate := schematemplate.New(targetModule.DirName())
	if err := schemaTemplate.Export(targetModule.AbsPath()); err != nil {
		return errors.Wrap(err, "failed to generate schema file")
	}

	return nil
}

func (a *api) replaceResolvers(targetModule module.Module) error {
	const resolversFilename = "schema.resolvers.go"
	resolversFile := cliio.FileOf(filepath.Join(targetModule.AbsPath(), schematemplate.DirName, resolversFilename))
	if err := resolversFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", resolversFile.AbsPath())
	}

	return nil
}
