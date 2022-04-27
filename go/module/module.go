package module

import (
	"github.com/funcgql/cli/shell"
)

type Module interface {
	Name() string
	DirName() string
	AbsPath() string
	Tidy(shellAPI shell.API) error
	ToolOf(importPath string) Tool
	Tools() ([]Tool, error)
	InstallInitialTools(shellAPI shell.API) error
	InstallAllTools(shellAPI shell.API) error
}

func (m module) Name() string {
	return m.name
}

func (m module) DirName() string {
	return m.dirName
}

func (m module) AbsPath() string {
	return m.absPath
}

type module struct {
	name    string
	dirName string
	absPath string
}
