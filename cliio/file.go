package cliio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type File interface {
	AbsPath() string
	StringContent() (string, error)
	Write(content string) error
	WriteBytes(bytes []byte) error
	Make() error
	Exists() (bool, error)
	Remove() error
}

func FileOf(absPath string) File {
	return &file{absPath: absPath}
}

func TempFile(name string) (File, error) {
	iofile, err := ioutil.TempFile("", name)
	if err != nil {
		return nil, err
	}

	return &file{absPath: iofile.Name()}, nil
}

func (f *file) AbsPath() string {
	return f.absPath
}

func (f *file) StringContent() (string, error) {
	bytes, err := ioutil.ReadFile(f.absPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read content of %s", f.AbsPath())
	}

	return string(bytes), nil
}

func (f *file) Write(content string) error {
	return f.WriteBytes([]byte(content))
}

func (f *file) WriteBytes(bytes []byte) error {
	iofile, err := f.make()
	if err != nil {
		return err
	}
	defer iofile.Close()

	if _, err = iofile.Write(bytes); err != nil {
		return err
	}

	return iofile.Sync()
}

func (f *file) Make() error {
	_, err := f.make()
	return err
}

func (f *file) Exists() (bool, error) {
	info, err := os.Stat(f.AbsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return !info.IsDir(), nil
}

func (f *file) Remove() error {
	return os.RemoveAll(f.absPath)
}

type file struct {
	absPath string
}

func (f *file) make() (*os.File, error) {
	// Create parent directory if needed.
	if err := DirOf(filepath.Dir(f.AbsPath())).Make(); err != nil {
		return nil, err
	}

	return os.Create(f.AbsPath())
}
