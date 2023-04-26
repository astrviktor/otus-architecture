package rotate

type RotateInterface interface {
	GetDirection() (int, error)
	SetDirection(n int) error
}

type RotateCommand struct {
	obj RotateInterface
}

func NewRotateCommand(obj RotateInterface) *RotateCommand {
	return &RotateCommand{obj: obj}
}

func (c *RotateCommand) Execute() error {
	n, err := c.obj.GetDirection()
	if err != nil {
		return err
	}

	n = n + 1

	return c.obj.SetDirection(n)
}
