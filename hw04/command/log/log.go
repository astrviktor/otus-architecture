package log

import (
	"otus-architecture/hw04/additions/logging"
)

type LogInterface interface {
	GetError() (error, error)
}

type LogCommand struct {
	obj LogInterface
}

func NewLogCommand(obj LogInterface) *LogCommand {
	return &LogCommand{obj: obj}
}

func (c *LogCommand) Execute() error {
	var e, err error

	if e, err = c.obj.GetError(); err != nil {
		return err
	}

	return logging.WriteErrorToLog(e)
}
