package log

import (
	"otus-architecture/hw11/additions/logging"
	"otus-architecture/hw11/object"
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

func (c *LogCommand) Execute() (*object.Object, error) {
	var e, err error

	if e, err = c.obj.GetError(); err != nil {
		return nil, err
	}

	return nil, logging.WriteErrorToLog(e)
}
