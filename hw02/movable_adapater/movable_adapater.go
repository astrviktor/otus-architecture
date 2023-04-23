package movable_adapater

import (
	"math"
	"otus-architecture/hw02/getters"
	"otus-architecture/hw02/uobject"
	"otus-architecture/hw02/vector"
)

type MovableAdapater struct {
	obj *uobject.UObject
}

func New(obj *uobject.UObject) MovableAdapater {
	return MovableAdapater{obj: obj}
}

func (ma *MovableAdapater) GetPosition() (vector.Vector, error) {
	return getters.GetVector(ma.obj.GetProperty("Position"))
}

func (ma *MovableAdapater) SetPosition(v vector.Vector) error {
	return ma.obj.SetProperty("Position", v)
}

func (a *MovableAdapater) GetVelocity() (vector.Vector, error) {
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
