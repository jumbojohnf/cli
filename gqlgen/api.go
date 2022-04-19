package gqlgen

import "github.com/funcgql/cli/shell"

type API interface {
	Init(absPath string) error

	execute(args ...string) (shell.Output, error)
}

func NewAPI() API {
	return &api{}
}

const moduleName = "github.com/99designs/gqlgen"

type api struct{}

func (a *api) execute(args ...string) (shell.Output, error) {
	goArgs := append([]string{"run", moduleName}, args...)
	return shell.Execute("go", goArgs...)
}
