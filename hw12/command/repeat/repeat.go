package repeat

import (
	"otus-architecture/hw12/command"
)

type RepeatCommand struct {
	cmd command.CommandInterface
}

func NewRepeatCommand(cmd command.CommandInterface) *RepeatCommand {
	return &RepeatCommand{cmd: cmd}
}

func (c *RepeatCommand) Execute() error {
	return c.cmd.Execute()
}
