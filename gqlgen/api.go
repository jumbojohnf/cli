package gqlgen

type API interface {
	Init(absPath string, moduleName string) error
}

func NewAPI() API {
	return &api{}
}

const gqlgenModuleName = "github.com/99designs/gqlgen"

type api struct{}
