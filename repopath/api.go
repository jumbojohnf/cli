package repopath

import (
	"strings"

	"github.com/funcgql/cli/shell"
)

type API interface {
	RootPath() (string, error)
}

func NewAPI(shellAPI shell.API) API {
	return &api{
		shellAPI: shellAPI,
	}
}

func (a *api) RootPath() (string, error) {
	output, err := a.shellAPI.Execute("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}
	return strings.Trim(output.Combined, " \n"), nil
}

type api struct {
	shellAPI shell.API
}
