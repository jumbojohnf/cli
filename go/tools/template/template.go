package template

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/template"
)

type GoToolsTemplate interface {
	Export(rootDir string) (*cliio.File, error)
}

func New(moduleDirName string) GoToolsTemplate {
	return goToolsTemplate{
		moduleDirName: moduleDirName,
	}
}

func (t goToolsTemplate) Export(dirAbsPath string) (*cliio.File, error) {
	content, err := t.render()
	if err != nil {
		return nil, err
	}

	const filename = "tools.go"
	return template.Export(content, filepath.Join(dirAbsPath, filename))
}

//go:embed tools.go.template
var toolsGoTemplate string

type goToolsTemplate struct {
	moduleDirName string
}

func (t goToolsTemplate) contentData() (interface{}, error) {
	type templateData struct {
		PackageName string
	}
	return templateData{
		PackageName: t.moduleDirName,
	}, nil
}

func (t goToolsTemplate) render() (string, error) {
	data, err := t.contentData()
	if err != nil {
		return "", err
	}
	return template.Render("toolsgo", toolsGoTemplate, data)
}
