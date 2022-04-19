package template

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/go/version"
	"github.com/funcgql/cli/template"
)

type GoModTemplate interface {
	Export(rootDir string) (*cliio.File, error)
}

func New(moduleName string) GoModTemplate {
	return goModTemplate{
		moduleName: moduleName,
	}
}

func (t goModTemplate) Export(dirAbsPath string) (*cliio.File, error) {
	content, err := t.render()
	if err != nil {
		return nil, err
	}

	const filename = "go.mod"
	return template.Export(content, filepath.Join(dirAbsPath, filename))
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
		GoVersion:  version.Current(),
	}, nil
}

func (t goModTemplate) render() (string, error) {
	data, err := t.templateContentData()
	if err != nil {
		return "", err
	}
	return template.Render("gomod", goModTemplateContent, data)
}
