package rotate

import "otus-architecture/hw11/object"

type RotateCommand struct {
	obj *RotateAdapter
}

func NewRotateCommand(obj *RotateAdapter) *RotateCommand {
	return &RotateCommand{obj: obj}
}

func (c *RotateCommand) GetObject() *object.Object {
	return c.obj.GetObject()
}

func (c *RotateCommand) Execute() (*object.Object, error) {
	n, err := c.obj.GetDirection()
	if err != nil {
		return nil, err
	}

	n = n + 1

	err = c.obj.SetDirection(n)

	return c.obj.obj, err
}
