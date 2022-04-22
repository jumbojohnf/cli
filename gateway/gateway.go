package gateway

import (
	"path/filepath"

	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
)

type Gateway interface {
	ExportIndexFile(cfg *config.Config) error
	InstallPackages() error
}

func New(functionType functype.FunctionType, cfg *config.Config) Gateway {
	return gateway{
		functionType: functionType,
		absPath:      filepath.Join(cfg.GraphModulesAbsPath, dirName),
	}
}

const dirName = "gateway"

type gateway struct {
	functionType functype.FunctionType
	absPath      string
}
