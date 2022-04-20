package gqlgen

import (
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
)

type API interface {
	Init(absPath string, targetModule module.Module, functionTypes []functype.FunctionType) error
	Generate(absPath string) error
}

func NewAPI() API {
	return &api{}
}

type api struct{}
