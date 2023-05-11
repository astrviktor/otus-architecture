package adapter

import (
	"fmt"
	"os/exec"
)

type ExecutorInterface interface {
	Execute(app, input, output string) error
}

type Executor struct{}

func (e *Executor) Execute(app, input, output string) error {
	args := []string{"-input", input, "-output", output}

	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	fmt.Println(string(stdout))

	return err
}
