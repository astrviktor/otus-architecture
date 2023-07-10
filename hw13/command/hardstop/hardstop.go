package hardstop

import "otus-architecture/hw13/object"

type HardStopInterface interface {
	GetHardStopChannel() (chan struct{}, error)
}

type HardStopCommand struct {
	obj HardStopInterface
}

func NewHardStopCommand(obj HardStopInterface) *HardStopCommand {
	return &HardStopCommand{obj: obj}
}

func (c *HardStopCommand) Execute() (*object.Object, error) {
	hardStop, err := c.obj.GetHardStopChannel()
	if err != nil {
		return nil, err
	}

	hardStop <- struct{}{}

	return nil, nil
}
