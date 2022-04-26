package gqlgen

import (
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/shell"
)

type API interface {
	Init(targetModule module.Module, functionTypes []functype.FunctionType) error
	Generate(targetModule module.Module) error
}

func NewAPI(shellAPI shell.API) API {
	return &api{
		shellAPI: shellAPI,
	}
}

type api struct {
	shellAPI shell.API
}
