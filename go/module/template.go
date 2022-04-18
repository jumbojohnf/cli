package module

import (
	_ "embed"
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/go/version"
	"github.com/funcgql/cli/template"
)

//go:embed go.mod.template
var goModTemplate string

func (m Module) NewTemplate() template.Template {
	return m
}

func (m Module) Render(_ string) (string, error) {
	data, err := m.templateContentData()
	if err != nil {
		return "", err
	}
	return template.Render("gomod", goModTemplate, data)
}

func (m Module) Export(rootDir string) (*cliio.File, error) {
	content, err := m.Render(rootDir)
	if err != nil {
		return nil, err
	}

	const filename = "go.mod"
	return template.Export(content, filepath.Join(rootDir, m.DirName(), filename))
}

func (m Module) templateContentData() (interface{}, error) {
	type templateData struct {
		ModuleName string
		GoVersion  string
	}
	return templateData{
		ModuleName: m.Name(),
		GoVersion:  version.Current(),
	}, nil
}
