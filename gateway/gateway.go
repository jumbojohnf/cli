package gateway

import (
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
)

type Gateway interface {
	ExportIndexFile(cfg *config.Config) error
}

func New(functionType functype.FunctionType) Gateway {
	return gateway{functionType: functionType}
}

type gateway struct {
	functionType functype.FunctionType
}
