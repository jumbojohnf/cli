package shell

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func (a *api) Execute(cmdName string, args ...string) (Output, error) {
	return a.ExecuteIn("", cmdName, args...)
}

func (a *api) ExecuteWithIO(cmdName string, args ...string) (Output, error) {
	return a.ExecuteWithIOIn("", cmdName, args...)
}

func (a *api) ExecuteIn(dir string, cmdName string, args ...string) (Output, error) {
	cmd, err := baseCommand(cmdName, args...)
	if err != nil {
		return Output{}, err
	}

	cmd.Dir = dir

	return execute(cmd)
}

func (a *api) ExecuteWithIOIn(dir string, cmdName string, args ...string) (Output, error) {
	cmd, err := baseCommand(cmdName, args...)
	if err != nil {
		return Output{}, err
	}

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return execute(cmd)
}

func baseCommand(cmdName string, args ...string) (*exec.Cmd, error) {
	cmdPath, err := exec.LookPath(cmdName)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(cmdPath, args...)
	cmd.Env = os.Environ()
	return cmd, nil
}

func execute(cmd *exec.Cmd) (Output, error) {
	// Setup buffers to capture output.
	outputStream := &bytes.Buffer{}
	errorStream := &bytes.Buffer{}
	if cmd.Stdout != nil {
		cmd.Stdout = io.MultiWriter(outputStream, cmd.Stdout)
	} else {
		cmd.Stdout = outputStream
	}
	if cmd.Stderr != nil {
		cmd.Stderr = io.MultiWriter(errorStream, cmd.Stderr)
	} else {
		cmd.Stderr = errorStream
	}

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
