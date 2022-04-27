package gateway

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/gateway/template/index/lambda"
	"github.com/pkg/errors"
)

func New(functionType functype.FunctionType, cfg *config.Config) (Gateway, error) {
	gatewayAbsPath := filepath.Join(cfg.GraphModulesAbsPath, dirName)
	if alreadyExists, err := cliio.DirOf(gatewayAbsPath).Exists(); err != nil {
		return nil, errors.Wrapf(err, "failed to determine if gateway already exists at %s", gatewayAbsPath)
	} else if alreadyExists {
		return nil, errors.Errorf("gateway already exists at %s", gatewayAbsPath)
	}

	if err := exportIndexFile(functionType, cfg, gatewayAbsPath); err != nil {
		return nil, err
	}

	return gateway{
		functionType: functionType,
		absPath:      gatewayAbsPath,
	}, nil
}

const dirName = "gateway"

func exportIndexFile(functionType functype.FunctionType, cfg *config.Config, gatewayAbsPath string) error {
	switch functionType {
	case functype.Lambda:
		return lambda.Export(gatewayAbsPath, cfg.AWS.Gateway)
	default:
		return errors.Errorf("unknown function type %s for exporting gateway index file", functionType)
	}
}
