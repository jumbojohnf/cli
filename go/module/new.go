package module

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	modtemplate "github.com/funcgql/cli/go/module/template"
	toolstemplate "github.com/funcgql/cli/go/tools/template"
	worktemplate "github.com/funcgql/cli/go/work/template"
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
	if _, err := toolsGoTemplate.Export(absPath); err != nil {
		return nil, errors.Wrap(err, "failed to create new module tools.go file")
	}

	modTemplate := modtemplate.New(name)
	if _, err := modTemplate.Export(absPath); err != nil {
		return nil, errors.Wrap(err, "failed to create new go.mod file")
	}

	// Export go.work file after module is created since it needs to include the newly created module.
	goWorkTemplate := worktemplate.New()
	if _, err := goWorkTemplate.Export(rootAbsPath); err != nil {
		return nil, errors.Wrapf(err, "failed to update go.work file in %s", rootAbsPath)
	}

	return module{
		name:    name,
		dirName: dirName,
		absPath: absPath,
	}, nil
}
