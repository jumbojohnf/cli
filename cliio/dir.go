package cliio

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Dir struct {
	absPath string
}

func DirOf(absPath string) *Dir {
	return &Dir{absPath: absPath}
}

func TempDir(name string) (*Dir, error) {
	// ioutil.TempDir will replace the `*` with a random string.
	dirName := fmt.Sprintf("*-%s", name)
	absPath, err := ioutil.TempDir("", dirName)
	if err != nil {
		return nil, err
	}

	return &Dir{absPath: absPath}, nil
}

func (d *Dir) AbsPath() string {
	return d.absPath
}

func (d *Dir) Make() error {
	return os.MkdirAll(d.AbsPath(), os.ModePerm)
}

func (d *Dir) Exists() (bool, error) {
	info, err := os.Stat(d.AbsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}
