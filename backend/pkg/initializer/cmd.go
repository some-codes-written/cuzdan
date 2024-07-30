package initializer

import (
	"os"
	"os/exec"
)

func Godoc(file string) {

	cmd := &exec.Cmd{
		Path:   file,
		Args:   []string{file, ""},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	cmd.Start()

	cmd.Wait()

}
