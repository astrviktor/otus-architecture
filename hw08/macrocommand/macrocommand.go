package macrocommand

import (
	"otus-architecture/hw07/command"
)

type MacroCommand struct {
	commands []command.CommandInterface
}

func NewMacroCommand(commands ...command.CommandInterface) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (c *MacroCommand) Execute() error {
	for _, cmd := range c.commands {
		if err := cmd.Execute(); err != nil {
			return err
		}
	}

	return nil
}
