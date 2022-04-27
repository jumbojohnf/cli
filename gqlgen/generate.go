package gqlgen

import "github.com/funcgql/cli/go/module"

func (a *api) Generate(targetModule module.Module) error {
	return a.run("generate", targetModule)
}
