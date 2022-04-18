package module

type Module struct {
	name    string
	dirName string
	absPath string
}

func WithNameAndAbsPath(name string, absPath string) Module {
	return Module{
		name:    name,
		dirName: DirNameFromModuleName(name),
		absPath: absPath,
	}
}

func (m Module) Name() string {
	return m.name
}

func (m Module) DirName() string {
	return m.dirName
}

func (m Module) AbsPath() string {
	return m.absPath
}
