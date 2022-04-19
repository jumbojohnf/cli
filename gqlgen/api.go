package gqlgen

import "github.com/funcgql/cli/functype"

type API interface {
	Init(absPath string, moduleName string, functionType functype.FunctionType) error
}

func NewAPI() API {
	return &api{}
}

const gqlgenModuleName = "github.com/99designs/gqlgen"

type api struct{}
