package repeatcount

import (
	"otus-architecture/hw12/command"
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

func (c *RepeatCountCommand) Execute() error {
	if c.repeat > c.count {
		return nil
	}

	c.repeat++
	return c.cmd.Execute()
}
