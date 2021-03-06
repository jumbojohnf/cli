package cliio

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Dir interface {
	AbsPath() string
	Make() error
	Exists() (bool, error)
}

func DirOf(absPath string) Dir {
	return &dir{absPath: absPath}
}

func TempDir(name string) (Dir, error) {
	// ioutil.TempDir will replace the `*` with a random string.
	dirName := fmt.Sprintf("*-%s", name)
	absPath, err := ioutil.TempDir("", dirName)
	if err != nil {
		return nil, err
	}

	return &dir{absPath: absPath}, nil
}

func (d *dir) AbsPath() string {
	return d.absPath
}

func (d *dir) Make() error {
	return os.MkdirAll(d.AbsPath(), os.ModePerm)
}

func (d *dir) Exists() (bool, error) {
	info, err := os.Stat(d.AbsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

type dir struct {
	absPath string
}
