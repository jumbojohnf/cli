package gateway

import (
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/gateway/template/index/lambda"
	"github.com/pkg/errors"
)

func (g gateway) ExportIndexFile(cfg *config.Config) error {
	switch g.functionType {
	case functype.Lambda:
		return lambda.Export(g.absPath, cfg.AWS.Gateway)
	default:
		return errors.Errorf("unknown function type %s for exporting gateway index file", g.functionType)
	}
}
