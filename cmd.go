package hooks

import (
	"os"
	"os/exec"
)

func NewCmd(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	return cmd
}

func NewCmdHook(name string, args ...string) func() error {
	return func() error {
		return NewCmd(name, args...).Run()
	}
}
