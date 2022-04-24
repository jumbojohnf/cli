package module

import (
	"github.com/funcgql/cli/go/tools"
)

func (m module) InstallTools(toolsAPI tools.API) error {
	return toolsAPI.InstallAllIn(m.absPath)
}
