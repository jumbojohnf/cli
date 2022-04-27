package gateway

import (
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/npm"
)

type Gateway interface {
	InstallPackages(npmAPI npm.API) error
}

type gateway struct {
	functionType functype.FunctionType
	absPath      string
}
