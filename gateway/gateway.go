package gateway

import (
	"path/filepath"

	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/npm"
)

type Gateway interface {
	ExportIndexFile(cfg *config.Config) error
	InstallPackages(npmAPI npm.API) error
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
