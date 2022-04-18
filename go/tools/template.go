package tools

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/template"
)

//go:embed tools.go.template
var toolsGoTemplate string

func NewTemplate(moduleName string) template.Template {
	moduleDirName := filepath.Base(moduleName)
	return goToolsTemplate{moduleDirName: moduleDirName}
}

type goToolsTemplate struct {
	moduleDirName string
}

func (t goToolsTemplate) Render(_ string) (string, error) {
	data, err := t.contentData()
	if err != nil {
		return "", err
	}
	return template.Render("toolsgo", toolsGoTemplate, data)
}

func (t goToolsTemplate) Export(rootDir string) (*cliio.File, error) {
	content, err := t.Render(rootDir)
	if err != nil {
		return nil, err
	}

	const filename = "tools.go"
	return template.Export(content, filepath.Join(rootDir, t.moduleDirName, filename))
}

func (t goToolsTemplate) contentData() (interface{}, error) {
	type templateData struct {
		PackageName string
	}
	return templateData{
		PackageName: t.moduleDirName,
	}, nil
}
