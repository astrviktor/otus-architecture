package hardstop

type HardStopInterface interface {
	GetHardStopChannel() (chan struct{}, error)
}

type HardStopCommand struct {
	obj HardStopInterface
}

func NewHardStopCommand(obj HardStopInterface) *HardStopCommand {
	return &HardStopCommand{obj: obj}
}

func (c *HardStopCommand) Execute() error {
	hardStop, err := c.obj.GetHardStopChannel()
	if err != nil {
		return err
	}

	hardStop <- struct{}{}

	return nil
}
