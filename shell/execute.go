package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type Output struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Combined string
}

func Execute(cmdName string, args ...string) (Output, error) {
	return ExecuteIn("", cmdName, args...)
}

func ExecuteIn(dir string, cmdName string, args ...string) (Output, error) {
	cmdPath, err := exec.LookPath(cmdName)
	if err != nil {
		return Output{}, err
	}

	outputStream := &bytes.Buffer{}
	errorStream := &bytes.Buffer{}
	cmd := exec.Command(cmdPath, args...)
	cmd.Dir = dir
	cmd.Env = os.Environ()
	cmd.Stdout = outputStream
	cmd.Stderr = errorStream

	executionErr := cmd.Run()
	stdout := strings.TrimSpace(outputStream.String())
	stderr := strings.TrimSpace(errorStream.String())
	combined := strings.TrimSpace(fmt.Sprintf("%s\n%s", stdout, stderr))

	output := Output{
		Stdout:   stdout,
		Stderr:   stderr,
		Combined: combined,
	}

	if executionErr != nil {
		if exitError, isExitError := executionErr.(*exec.ExitError); isExitError {
			output.ExitCode = exitError.ExitCode()
		}
		return output, errors.Wrap(executionErr, combined)
	}

	return output, nil
}
