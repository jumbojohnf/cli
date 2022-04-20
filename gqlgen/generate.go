package gqlgen

func (a *api) Generate(absPath string) error {
	return a.runIn("generate", absPath)
}
