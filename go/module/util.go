package module

import (
	"path/filepath"

	"github.com/funcgql/cli/shell"
)

func DirNameFromModuleName(moduleName string) string {
	return filepath.Base(moduleName)
}

func (m Module) execute(cmd string, args ...string) (shell.Output, error) {
	cmdArgs := append([]string{"mod", cmd}, args...)
	return shell.ExecuteIn(m.absPath, "go", cmdArgs...)
}
