package gomod

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/go/version"
	"github.com/funcgql/cli/template"
)

const Filename = "go.mod"

type GoModTemplate interface {
	Export(rootDir string) error
}

func New(moduleName string) GoModTemplate {
	return goModTemplate{
		moduleName: moduleName,
	}
}

func (t goModTemplate) Export(dirAbsPath string) error {
	content, err := t.render()
	if err != nil {
		return err
	}

	if _, err := template.Export(content, filepath.Join(dirAbsPath, Filename)); err != nil {
		return err
	}
	return nil
}

//go:embed go.mod.template
var goModTemplateContent string

type goModTemplate struct {
	moduleName string
}

func (t goModTemplate) templateContentData() (interface{}, error) {
	type templateData struct {
		ModuleName string
		GoVersion  string
	}
	return templateData{
		ModuleName: t.moduleName,
		GoVersion:  version.Current().MajorMinor(),
	}, nil
}

func (t goModTemplate) render() (string, error) {
	data, err := t.templateContentData()
	if err != nil {
		return "", err
	}
	return template.Render("gomod", goModTemplateContent, data)
}
