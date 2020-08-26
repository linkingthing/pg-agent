package util

import (
	"os"
	"os/exec"
)

func ExecCommand(args string) error {
	cmd := exec.Command("/bin/sh", "-c", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
