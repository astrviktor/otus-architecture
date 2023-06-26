package move

import (
	"errors"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/object"
)

type MoveCommand struct {
	obj *MoveAdapter
}

func NewMoveCommand(obj *MoveAdapter) *MoveCommand {
	return &MoveCommand{obj: obj}
}

func (c *MoveCommand) Execute() (*object.Object, error) {
	var position, velocity vector.Vector
	var err error

	if position, err = c.obj.GetPosition(); err != nil {
		return nil, err
	}

	if velocity, err = c.obj.GetVelocity(); err != nil {
		return nil, err
	}

	err = c.obj.SetPosition(vector.VectorPlus(position, velocity))
	return c.obj.Object, err
}

var (
	ErrPositionFreeze = errors.New("position freeze")
)
