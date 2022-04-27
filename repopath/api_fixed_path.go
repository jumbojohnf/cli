package repopath

func FixedPathAPI(rootPath string) API {
	return &fixedPathAPI{rootPath: rootPath}
}

func (a *fixedPathAPI) RootPath() (string, error) {
	return a.rootPath, nil
}

type fixedPathAPI struct {
	rootPath string
}
