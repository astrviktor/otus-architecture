package run

type RunInterface interface {
	GetRunChannel() (chan struct{}, error)
}

type RunCommand struct {
	obj RunInterface
}

func NewRunCommand(obj RunInterface) *RunCommand {
	return &RunCommand{obj: obj}
}

func (c *RunCommand) Execute() error {
	run, err := c.obj.GetRunChannel()
	if err != nil {
		return err
	}

	close(run)

	return nil
}
