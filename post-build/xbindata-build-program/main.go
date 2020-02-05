package xbindata_build_program

import (
	"os"
	"os/exec"
)

func Run(binaries ...string) (err error) {
	if len(binaries) == 0 {
		return
	}
	cmd := exec.Command("xbindata", append([]string{"build", "program"}, binaries...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	return cmd.Run()
}
