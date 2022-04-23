package module

import "github.com/funcgql/cli/go/tools"

func (m module) InstallTools() error {
	return tools.InstallAllIn(m.absPath)
}
