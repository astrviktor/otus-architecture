package repeatcount

import (
	"otus-architecture/hw13/command"
	"otus-architecture/hw13/object"
)

type RepeatCountCommand struct {
	repeat int
	count  int
	cmd    command.CommandInterface
}

func NewRepeatCountCommand(cmd command.CommandInterface, count int) *RepeatCountCommand {
	return &RepeatCountCommand{
		repeat: 0,
		count:  count,
		cmd:    cmd,
	}
}

func (c *RepeatCountCommand) Execute() (*object.Object, error) {
	if c.repeat > c.count {
		return nil, nil
	}

	c.repeat++
	return c.cmd.Execute()
}
