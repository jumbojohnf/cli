package module

import (
	"os"
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

func WithName(name string, cfg *config.Config) (Module, bool, error) {
	dirName := filepath.Base(name)
	absPath := filepath.Join(cfg.GraphModulesAbsPath, dirName)

	if dirExists, err := cliio.DirOf(absPath).Exists(); err != nil {
		return nil, false, errors.Wrapf(err, "failed to determine if module %s directory %s exists", name, absPath)
	} else if !dirExists {
		return nil, false, nil
	}

	return module{
		name:    name,
		dirName: dirName,
		absPath: absPath,
	}, true, nil
}

func CurrentDir(shellAPI shell.API) (Module, bool, error) {
	workingDirPath, err := os.Getwd()
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to determine current working directory path")
	}

	dirName := filepath.Base(workingDirPath)

	output, err := shellAPI.ExecuteIn(workingDirPath, "go", "list", "-m")
	if err != nil {
		return nil, false, errors.Wrapf(err, "failed to obtain module name in %s", workingDirPath)
	}

	return module{
		name:    output.Combined,
		dirName: dirName,
		absPath: workingDirPath,
	}, true, nil
}
