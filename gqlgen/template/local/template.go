package local

import (
	_ "embed"

	"github.com/funcgql/cli/gqlgen/template/commonmain"
)

type MainTemplate interface {
	Export(rootDir string) error
}

func New(moduleName string) MainTemplate {
	return mainTemplate{
		commonmain.New(moduleName, templateContent),
	}
}

func (t mainTemplate) Export(rootDir string) error {
	const dirName = "local"
	return t.mainTemplate.Export(rootDir, dirName)
}

//go:embed main.go.template
var templateContent string

type mainTemplate struct {
	mainTemplate commonmain.MainTemplate
}
