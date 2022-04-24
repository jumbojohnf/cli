package gqlgen

import (
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/go/tools"
)

type API interface {
	Init(absPath string, targetModule module.Module, functionTypes []functype.FunctionType) error
	Generate(absPath string) error
}

func NewAPI(toolsAPI tools.API) API {
	return &api{
		toolsAPI: toolsAPI,
	}
}

type api struct {
	toolsAPI tools.API
}
