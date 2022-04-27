package schema

import (
	_ "embed"
	"path/filepath"
	"strings"

	"github.com/funcgql/cli/template"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	DirName  = "graph"
	Filename = "schema.graphqls"
)

type SchemaTemplate interface {
	Export(rootDir string) error
}

func New(moduleDirName string) SchemaTemplate {
	return schemaTemplate{
		moduleDirName: moduleDirName,
	}
}

func (t schemaTemplate) Export(rootDir string) error {
	content, err := t.render(rootDir)
	if err != nil {
		return err
	}

	if _, err := template.Export(content, filepath.Join(rootDir, DirName, Filename)); err != nil {
		return err
	}
	return nil
}

//go:embed schema.graphqls.template
var schemaTemplateContent string

type schemaTemplate struct {
	moduleDirName string
}

func (t schemaTemplate) contentData() interface{} {
	type templateData struct {
		ModuleTypeName    string
		QueryFunctionName string
	}

	return templateData{
		ModuleTypeName:    cases.Title(language.English).String(t.moduleDirName),
		QueryFunctionName: strings.ToLower(t.moduleDirName),
	}
}

func (t schemaTemplate) render(rootDir string) (string, error) {
	data := t.contentData()
	return template.Render("schema", schemaTemplateContent, data)
}
