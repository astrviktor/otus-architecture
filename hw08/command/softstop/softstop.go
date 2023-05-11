package softstop

type SoftStopInterface interface {
	GetSoftStopChannel() (chan struct{}, error)
}

type SoftStopCommand struct {
	obj SoftStopInterface
}

func NewSoftStopCommand(obj SoftStopInterface) *SoftStopCommand {
	return &SoftStopCommand{obj: obj}
}

func (c *SoftStopCommand) Execute() error {
	softStop, err := c.obj.GetSoftStopChannel()
	if err != nil {
		return err
	}

	softStop <- struct{}{}

	return nil
}
