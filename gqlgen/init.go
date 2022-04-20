package gqlgen

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
	lambdatemplate "github.com/funcgql/cli/gqlgen/template/lambda"
	schematemplate "github.com/funcgql/cli/gqlgen/template/schema"
	"github.com/pkg/errors"
)

func (a *api) Init(absPath string, targetModule module.Module, functionTypes []functype.FunctionType) error {
	if err := a.runIn("init", absPath); err != nil {
		return err
	}

	if err := a.replaceMain(absPath, targetModule, functionTypes); err != nil {
		return err
	}
	if err := a.replaceSchema(absPath, targetModule); err != nil {
		return err
	}
	if err := a.replaceResolvers(absPath); err != nil {
		return err
	}

	return nil
}

func (a *api) replaceMain(absPath string, targetModule module.Module, functionTypes []functype.FunctionType) error {
	serverFile := cliio.FileOf(filepath.Join(absPath, "server.go"))
	if err := serverFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", serverFile.AbsPath())
	}

	for _, functionType := range functionTypes {
		switch functionType {
		case functype.Lambda:
			mainTemplate := lambdatemplate.New(targetModule.Name())
			if err := mainTemplate.Export(absPath); err != nil {
				return errors.Wrap(err, "failed to generate main.go")
			}
		default:
			return errors.Errorf("unknown function type %s to generate main.go", functionType)
		}
	}

	return nil
}

func (a *api) replaceSchema(absPath string, targetModule module.Module) error {
	const (
		graphDirName   = "graph"
		schemaFilename = "schema.graphqls"
	)
	schemaFile := cliio.FileOf(filepath.Join(absPath, graphDirName, schemaFilename))
	if err := schemaFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", schemaFile.AbsPath())
	}

	schemaTemplate := schematemplate.New(targetModule.DirName())
	if err := schemaTemplate.Export(absPath); err != nil {
		return errors.Wrap(err, "failed to generate schema file")
	}

	return nil
}

func (a *api) replaceResolvers(absPath string) error {
	const (
		graphDirName      = "graph"
		resolversFilename = "schema.resolvers.go"
	)
	resolversFile := cliio.FileOf(filepath.Join(absPath, graphDirName, resolversFilename))
	if err := resolversFile.Remove(); err != nil {
		return errors.Wrapf(err, "failed to remove %s", resolversFile.AbsPath())
	}

	return nil
}
