package move

import (
	"errors"
	"otus-architecture/hw07/additions/vector"
)

type MoveInterface interface {
	GetPosition() (vector.Vector, error)
	SetPosition(v vector.Vector) error
	GetVelocity() (vector.Vector, error)
}

type MoveCommand struct {
	obj MoveInterface
}

func NewMoveCommand(obj MoveInterface) *MoveCommand {
	return &MoveCommand{obj: obj}
}

func (c *MoveCommand) Execute() error {
	var position, velocity vector.Vector
	var err error

	if position, err = c.obj.GetPosition(); err != nil {
		return err
	}

	if velocity, err = c.obj.GetVelocity(); err != nil {
		return err
	}

	return c.obj.SetPosition(vector.VectorPlus(position, velocity))
}

var (
	ErrPositionFreeze = errors.New("position freeze")
)
