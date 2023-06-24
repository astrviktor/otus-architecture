package rotate

import "otus-architecture/hw12/object"

type RotateInterface interface {
	GetObject() *object.Object
	GetDirection() (int, error)
	SetDirection(n int) error
}

type RotateCommand struct {
	obj RotateInterface
}

func NewRotateCommand(obj RotateInterface) *RotateCommand {
	return &RotateCommand{obj: obj}
}

func (c *RotateCommand) GetObject() *object.Object {
	return c.obj.GetObject()
}

func (c *RotateCommand) Execute() error {
	n, err := c.obj.GetDirection()
	if err != nil {
		return err
	}

	n = n + 1

	return c.obj.SetDirection(n)
}
