package config

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/funcgql/cli/repopath"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const ConfigFilename = "funcgql.yaml"

type Config struct {
	GraphModulesRelPath string `yaml:"graphModulesRelPath"`
	GraphModulesAbsPath string
	AWS                 *AWSConfig `yaml:"aws,omitempty"`
}

func LoadFromRepoRoot(repoPathAPI repopath.API) (*Config, error) {
	repoRoot, err := repoPathAPI.RootPath()
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine Git repository path")
	}

	result, hasConfig, err := LoadFrom(repoRoot)
	if err != nil {
		return nil, err
	} else if !hasConfig {
		return nil, errors.Errorf("cannot locate %s config file in %s", ConfigFilename, repoRoot)
	}
	return result, nil
}

func LoadFrom(dir string) (*Config, bool, error) {
	configFilePath, err := configFilePathIn(dir)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to find configuration file")
	}
	if len(configFilePath) <= 0 {
		return nil, false, nil
	}

	configContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, false, errors.Wrapf(err, "failed to read configuration file at %s", configFilePath)
	}
	var result Config
	if err := yaml.Unmarshal(configContent, &result); err != nil {
		return nil, false, errors.Wrapf(err, "failed to unmarshal configuration file at %s", configFilePath)
	}
	result.GraphModulesAbsPath = filepath.Join(dir, result.GraphModulesRelPath)
	return &result, true, nil
}

func configFilePathIn(dir string) (string, error) {
	var configFilePath string
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
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
