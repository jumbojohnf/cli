package npm

import "github.com/funcgql/cli/shell"

type API interface {
	InstallIn(dir string, packages ...string) error
}

func NewAPI(shellAPI shell.API) API {
	return &api{
		shellAPI: shellAPI,
	}
}

type api struct {
	shellAPI shell.API
}
