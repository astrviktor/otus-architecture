package takeobject

import (
	"errors"
	"otus-architecture/hw08/object"
)

type TakeObjectCommand struct {
	obj *TakeObjectAdapter
}

func NewTakeObjectCommand(obj *TakeObjectAdapter) *TakeObjectCommand {
	return &TakeObjectCommand{obj: obj}
}

func (c *TakeObjectCommand) Execute() (*object.Object, error) {
	gameIdentifier, err := c.obj.GetGameIdentifier()
	if err != nil {
		return nil, err
	}

	objectIdentifier, err := c.obj.GetObjectIdentifier()
	if err != nil {
		return nil, err
	}

	game, ok := object.Games[gameIdentifier]
	if !ok {
		return nil, ErrGameNotFound
	}

	obj, ok := game[objectIdentifier]
	if !ok {
		return nil, ErrObjectNotFound
	}

	return obj, nil
}

var (
	ErrGameNotFound   = errors.New("game not found")
	ErrObjectNotFound = errors.New("object not found")
)
