package move

import (
	"math"
	"otus-architecture/hw07/additions/getters"
	"otus-architecture/hw07/additions/vector"
	"otus-architecture/hw07/object"
)

type MoveAdapterMock struct {
	*object.Object
}

func NewMoveAdapterMock(obj *object.Object) *MoveAdapterMock {
	return &MoveAdapterMock{obj}
}

func (a *MoveAdapterMock) GetPosition() (vector.Vector, error) {
	return getters.GetVector(a.GetProperty("Position"))
}

func (a *MoveAdapterMock) SetPosition(v vector.Vector) error {
	return ErrPositionFreeze
}

func (a *MoveAdapterMock) GetVelocity() (vector.Vector, error) {
	d, err := getters.GetInt(a.GetProperty("Direction"))
	if err != nil {
		return vector.Vector{}, err
	}

	n, err := getters.GetInt(a.GetProperty("DirectionNumber"))
	if err != nil {
		return vector.Vector{}, err
	}

	v, err := getters.GetInt(a.GetProperty("Velocity"))
	if err != nil {
		return vector.Vector{}, err
	}

	dx := int(float64(v) * math.Cos(2*math.Pi*float64(n)/float64(d)))
	dy := int(float64(v) * math.Sin(2*math.Pi*float64(n)/float64(d)))

	velocity := vector.Vector{
		X: dx,
		Y: dy,
	}

	return velocity, nil
}
