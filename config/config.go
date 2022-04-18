package config

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/funcgql/cli/repopath"
	"github.com/pkg/errors"
)

const ConfigFilename = "funcgql.json"

type Config struct {
	RootPath    string `json:"rootPath"`
	RootAbsPath string
	AWS         *AWSConfig `json:"aws,omitempty"`
}

func LoadFromRepoRoot() (*Config, error) {
	repoRoot, err := repopath.GitRepoPath()
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine Git repository path")
	}

	return LoadFrom(repoRoot.Path)
}

func LoadFrom(rootPath string) (*Config, error) {
	configFilePath, err := configFilePathIn(rootPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find configuration file")
	}
	if len(configFilePath) <= 0 {
		return nil, nil
	}

	configContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read configuration file at %s", configFilePath)
	}
	var result Config
	if err := json.Unmarshal(configContent, &result); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal configuration file at %s", configFilePath)
	}
	result.RootAbsPath = filepath.Join(rootPath, result.RootPath)
	return &result, nil
}

func configFilePathIn(rootPath string) (string, error) {
	var configFilePath string
	if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if len(configFilePath) > 0 {
			return filepath.SkipDir
		}
		if d != nil && d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ConfigFilename) {
			configFilePath = path
		}
		return nil
	}); err != nil {
		return "", err
	}

	return configFilePath, nil
}
