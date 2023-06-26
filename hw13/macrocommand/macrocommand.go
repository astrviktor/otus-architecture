package macrocommand

import (
	"otus-architecture/hw13/command"
	"otus-architecture/hw13/object"
)

type MacroCommand struct {
	commands []command.CommandInterface
}

func NewMacroCommand(commands ...command.CommandInterface) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (c *MacroCommand) Execute() (*object.Object, error) {
	for _, cmd := range c.commands {
		if _, err := cmd.Execute(); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
