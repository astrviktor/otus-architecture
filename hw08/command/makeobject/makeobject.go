package makeobject

import (
	"otus-architecture/hw08/object"
)

type MakeObjectCommand struct {
	obj *MakeObjectAdapter
}

func NewMakeObjectCommand(obj *MakeObjectAdapter) *MakeObjectCommand {
	return &MakeObjectCommand{obj: obj}
}

func (c *MakeObjectCommand) Execute() (*object.Object, error) {
	gameIdentifier, err := c.obj.GetGameIdentifier()
	if err != nil {
		return nil, err
	}

	objectIdentifier, err := c.obj.GetObjectIdentifier()
	if err != nil {
		return nil, err
	}

	_, ok := object.Games[gameIdentifier]
	if !ok {
		object.Games[gameIdentifier] = make(map[int]*object.Object)
	}

	object.Games[gameIdentifier][objectIdentifier] = c.obj.Object
	//object.GameObjects[identifier] = c.obj.Object

	return c.obj.Object, nil
}
