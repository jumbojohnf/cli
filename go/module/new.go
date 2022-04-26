package module

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	modtemplate "github.com/funcgql/cli/go/module/template/gomod"
	toolstemplate "github.com/funcgql/cli/go/module/template/tools"
	"github.com/pkg/errors"
)

func New(name string, cfg *config.Config) (Module, error) {
	dirName := filepath.Base(name)
	absPath := filepath.Join(cfg.GraphModulesAbsPath, dirName)

	newModuleDir := cliio.DirOf(absPath)
	if alreadyExists, err := newModuleDir.Exists(); err != nil {
		return nil, errors.Wrapf(err, "failed to determine if new module directory %s already exists", newModuleDir.AbsPath())
	} else if alreadyExists {
		return nil, errors.Errorf("new module directory %s already exists", newModuleDir.AbsPath())
	}
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
