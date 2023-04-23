package hw02

import "otus-architecture/hw02/vector"

type MoveInterface interface {
	GetPosition() (vector.Vector, error)
	SetPosition(v vector.Vector) error
	GetVelocity() (vector.Vector, error)
}

type Move struct{}

func (move *Move) Execute(obj MoveInterface) error {
	var position, velocity vector.Vector
	var err error

	if position, err = obj.GetPosition(); err != nil {
		return err
	}

	if velocity, err = obj.GetVelocity(); err != nil {
		return err
	}

	return obj.SetPosition(vector.VectorPlus(position, velocity))
}
