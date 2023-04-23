package movable_adapater_mock

import (
	"math"
	"otus-architecture/hw02/getters"
	"otus-architecture/hw02/uobject"
	"otus-architecture/hw02/vector"
)

type MovableAdapaterMock struct {
	obj *uobject.UObject
}

func New(obj *uobject.UObject) MovableAdapaterMock {
	return MovableAdapaterMock{obj: obj}
}

func (ma *MovableAdapaterMock) GetPosition() (vector.Vector, error) {
	return getters.GetVector(ma.obj.GetProperty("Position"))
}

func (ma *MovableAdapaterMock) SetPosition(v vector.Vector) error {
	return nil
}

func (a *MovableAdapaterMock) GetVelocity() (vector.Vector, error) {
	d, err := getters.GetInt(a.obj.GetProperty("Direction"))
	if err != nil {
		return vector.Vector{}, err
	}

	n, err := getters.GetInt(a.obj.GetProperty("DirectionNumber"))
	if err != nil {
		return vector.Vector{}, err
	}

	v, err := getters.GetInt(a.obj.GetProperty("Velocity"))
	if err != nil {
		return vector.Vector{}, err
	}

	x := int(float64(v) * math.Cos(2*math.Pi*float64(n)/float64(d)))
	y := int(float64(v) * math.Sin(2*math.Pi*float64(n)/float64(d)))

	velocity := vector.Vector{
		X: x,
		Y: y,
	}

	return velocity, nil
}
