package move

import (
	"math"
	"otus-architecture/hw11/additions/getters"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/object"
)

type MoveAdapter struct {
	*object.Object
}

func NewMoveAdapter(obj *object.Object) *MoveAdapter {
	return &MoveAdapter{obj}
}

func (a *MoveAdapter) GetPosition() (vector.Vector, error) {
	return getters.GetVector(a.GetProperty("Position"))
}

func (a *MoveAdapter) SetPosition(v vector.Vector) error {
	a.SetProperty("Position", v)
	return nil
}

func (a *MoveAdapter) GetVelocity() (vector.Vector, error) {
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

	a.SetProperty("VelocityVector", velocity)

	return velocity, nil
}
