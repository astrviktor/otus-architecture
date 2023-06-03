package repeat

import (
	"otus-architecture/hw08/command"
	"otus-architecture/hw08/object"
)

type RepeatCommand struct {
	cmd command.CommandInterface
}

func NewRepeatCommand(cmd command.CommandInterface) *RepeatCommand {
	return &RepeatCommand{cmd: cmd}
}

func (c *RepeatCommand) Execute() (*object.Object, error) {
	return c.cmd.Execute()
}
