package module

import (
	"github.com/funcgql/cli/shell"
)

func (m module) execute(cmd string, args ...string) (shell.Output, error) {
	cmdArgs := append([]string{"mod", cmd}, args...)
	return shell.ExecuteIn(m.absPath, "go", cmdArgs...)
}
