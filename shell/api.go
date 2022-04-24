package shell

type API interface {
	Execute(cmdName string, args ...string) (Output, error)
	ExecuteIn(dir string, cmdName string, args ...string) (Output, error)
}

type Output struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Combined string
}

func NewAPI() API {
	return &api{}
}

type api struct{}
