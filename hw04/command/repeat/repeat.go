package repeat

import (
	"otus-architecture/hw04/command"
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
