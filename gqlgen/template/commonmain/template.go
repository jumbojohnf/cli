package commonmain

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/template"
)

type MainTemplate interface {
	Export(rootDir string, typeName string) error
}

func New(moduleName string, templateContent string) MainTemplate {
	return mainTemplate{
		moduleName:      moduleName,
		templateContent: templateContent,
	}
}

func (t mainTemplate) Export(rootDir string, typeName string) error {
	content, err := t.render(rootDir)
	if err != nil {
		return err
	}

	const filename = "main.go"
	if _, err := template.Export(content, filepath.Join(rootDir, typeName, filename)); err != nil {
		return err
	}
	return nil
}

type mainTemplate struct {
	moduleName      string
	templateContent string
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
	return template.Render("main", t.templateContent, data)
}
