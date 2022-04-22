package module

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	modtemplate "github.com/funcgql/cli/go/module/template"
	toolstemplate "github.com/funcgql/cli/go/tools/template"
	"github.com/pkg/errors"
)

func New(name string, rootAbsPath string) (Module, error) {
	dirName := filepath.Base(name)
	absPath := filepath.Join(rootAbsPath, dirName)

	newModuleDir := cliio.DirOf(absPath)
	if err := newModuleDir.Make(); err != nil {
		return nil, errors.Wrapf(err, "failed to create new module directory %s", absPath)
	}

	toolsGoTemplate := toolstemplate.New(dirName)
	if err := toolsGoTemplate.Export(absPath); err != nil {
		return nil, errors.Wrap(err, "failed to create new module tools.go file")
	}

	modTemplate := modtemplate.New(name)
	if err := modTemplate.Export(absPath); err != nil {
		return nil, errors.Wrap(err, "failed to create new go.mod file")
	}

	return module{
		name:    name,
		dirName: dirName,
		absPath: absPath,
	}, nil
}
