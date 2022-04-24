package tools

import "github.com/funcgql/cli/shell"

type API interface {
	InstallAllIn(dir string) error
	RunIn(moduleName string, absPath string, args ...string) (shell.Output, error)
}

func NewAPI(shellAPI shell.API) API {
	return &api{
		shellAPI: shellAPI,
	}
}

type api struct {
	shellAPI shell.API
}
