package tools

import "github.com/funcgql/cli/shell"

func RunIn(moduleName string, absPath string, args ...string) (shell.Output, error) {
	goArgs := append([]string{"run", moduleName}, args...)
	return shell.ExecuteIn(absPath, "go", goArgs...)
}
