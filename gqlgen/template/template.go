package template

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/template"
)

type MainTemplate interface {
	Export(rootDir string) error
}

func New(moduleName string) MainTemplate {
	return mainTemplate{
		moduleName: moduleName,
	}
}

func (t mainTemplate) Export(rootDir string) error {
	content, err := t.render(rootDir)
	if err != nil {
		return err
	}

	const filename = "main.go"
	if _, err := template.Export(content, filepath.Join(rootDir, filename)); err != nil {
		return err
	}
	return nil
}

//go:embed main.go.template
var mainTemplateContent string

type mainTemplate struct {
	moduleName string
}

func (t mainTemplate) contentData() interface{} {
	type templateData struct {
		ModuleName string
	}

	return templateData{
		ModuleName: t.moduleName,
	}
}

func (t mainTemplate) render(rootDir string) (string, error) {
	data := t.contentData()
	return template.Render("main", mainTemplateContent, data)
}
