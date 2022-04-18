package work

import (
	_ "embed"
	"io/fs"
	"path/filepath"
	"sort"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/go/version"
	"github.com/funcgql/cli/template"
	"github.com/pkg/errors"
)

//go:embed go.work.template
var goWorkTemplateContent string

func NewTemplate() template.Template {
	return goWorkTemplate{}
}

type goWorkTemplate struct{}

func (t goWorkTemplate) Render(rootDir string) (string, error) {
	data, err := t.contentData(rootDir)
	if err != nil {
		return "", err
	}
	return template.Render("gowork", goWorkTemplateContent, data)
}

func (t goWorkTemplate) Export(rootDir string) (*cliio.File, error) {
	content, err := t.Render(rootDir)
	if err != nil {
		return nil, err
	}

	const filename = "go.work"
	return template.Export(content, filepath.Join(rootDir, filename))
}

func (t goWorkTemplate) contentData(rootDir string) (interface{}, error) {
	type templateData struct {
		GoVersion      string
		ModuleDirNames []string
	}

	moduleDirNames, err := t.topLevelModuleDirNames(rootDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to obtain top-level module directory names in %s", rootDir)
	}
	return templateData{
		GoVersion:      version.Current(),
		ModuleDirNames: moduleDirNames,
	}, nil
}

func (t goWorkTemplate) topLevelModuleDirNames(rootPath string) ([]string, error) {
	var topLevelDirs []string
	if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d != nil && d.IsDir() && path != rootPath {
			topLevelDirs = append(topLevelDirs, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	var results []string
	for _, topLevelDir := range topLevelDirs {
		if err := filepath.WalkDir(topLevelDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if filepath.Base(path) == "go.mod" {
				dir := filepath.Dir(path)
				results = append(results, filepath.Base(dir))
				// Once found the top-level module, ignore the rest of the directory.
				return filepath.SkipDir
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results, nil
}
