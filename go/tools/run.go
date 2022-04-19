package tools

import "github.com/funcgql/cli/shell"

func Run(moduleName string, args ...string) (shell.Output, error) {
	goArgs := append([]string{"run", moduleName}, args...)
	return shell.Execute("go", goArgs...)
}
