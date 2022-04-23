package module

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
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
