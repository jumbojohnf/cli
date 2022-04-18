package cliio

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type File struct {
	absPath string
}

func FileOf(absPath string) *File {
	return &File{absPath: absPath}
}

func TempFile(name string) (*File, error) {
	iofile, err := ioutil.TempFile("", name)
	if err != nil {
		return nil, err
	}

	return &File{absPath: iofile.Name()}, nil
}

func (f *File) AbsPath() string {
	return f.absPath
}

func (f *File) StringContent() (string, error) {
	bytes, err := ioutil.ReadFile(f.absPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read content of %s", f.AbsPath())
	}

	return string(bytes), nil
}

func (f *File) Write(content string) error {
	return f.WriteBytes([]byte(content))
}

func (f *File) WriteBytes(bytes []byte) error {
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

func (f *File) Make() error {
	_, err := f.make()
	return err
}

func (f *File) Exists() (bool, error) {
	info, err := os.Stat(f.AbsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return !info.IsDir(), nil
}

func (f *File) make() (*os.File, error) {
	// Create parent directory if needed.
	if err := DirOf(filepath.Dir(f.AbsPath())).Make(); err != nil {
		return nil, err
	}

	return os.Create(f.AbsPath())
}
