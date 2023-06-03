package softstop

import "otus-architecture/hw08/object"

type SoftStopInterface interface {
	GetSoftStopChannel() (chan struct{}, error)
}

type SoftStopCommand struct {
	obj SoftStopInterface
}

func NewSoftStopCommand(obj SoftStopInterface) *SoftStopCommand {
	return &SoftStopCommand{obj: obj}
}

func (c *SoftStopCommand) Execute() (*object.Object, error) {
	softStop, err := c.obj.GetSoftStopChannel()
	if err != nil {
		return nil, err
	}

	softStop <- struct{}{}

	return nil, nil
}
