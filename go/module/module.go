package module

type Module interface {
	Name() string
	DirName() string
	AbsPath() string
	Tidy() error
}

func (m module) Name() string {
	return m.name
}

func (m module) DirName() string {
	return m.dirName
}

func (m module) AbsPath() string {
	return m.absPath
}

type module struct {
	name    string
	dirName string
	absPath string
}
