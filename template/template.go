package template

import (
	"bytes"
	"text/template"

	"github.com/funcgql/cli/cliio"
)

type Template interface {
	Render(rootDir string) (string, error)
	Export(dir string) (*cliio.File, error)
}

func Render(templateName string, templateContent string, data interface{}) (string, error) {
	temp, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer
	err = temp.Execute(&outputBuffer, data)
	if err != nil {
		return "", err
	}
	return outputBuffer.String(), nil
}

func Export(content string, filePath string) (cliio.File, error) {
	dstFile := cliio.FileOf(filePath)
	if err := dstFile.Make(); err != nil {
		return nil, err
	}

	if err := dstFile.Write(content); err != nil {
		return nil, err
	}
	return dstFile, nil
}
