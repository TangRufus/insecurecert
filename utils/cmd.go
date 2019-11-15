package utils

import (
	"os"
	"os/exec"
)

type Cmd struct {}

func (Cmd) Run(arg string, args ...string) error {
	cmd := exec.Command(arg, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
