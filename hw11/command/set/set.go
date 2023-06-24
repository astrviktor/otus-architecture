package set

import "otus-architecture/hw11/object"

type SetCommand struct {
	obj *SetAdapter
}

func NewSetCommand(obj *SetAdapter) *SetCommand {
	return &SetCommand{obj: obj}
}

func (c *SetCommand) Execute() (*object.Object, error) {
	return c.obj.Object, nil
}
