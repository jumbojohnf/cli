package tools

import "github.com/funcgql/cli/shell"

func (a *api) RunIn(moduleName string, absPath string, args ...string) (shell.Output, error) {
	goArgs := append([]string{"run", moduleName}, args...)
	return a.shellAPI.ExecuteIn(absPath, "go", goArgs...)
}
