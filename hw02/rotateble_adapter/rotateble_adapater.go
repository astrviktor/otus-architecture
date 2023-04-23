package rotateble_adapater

import (
	"otus-architecture/hw02/getters"
	"otus-architecture/hw02/uobject"
)

type RotatebleAdapater struct {
	obj *uobject.UObject
}

func New(obj *uobject.UObject) RotatebleAdapater {
	return RotatebleAdapater{obj: obj}
}

func (ma *RotatebleAdapater) GetDirection() (int, error) {
	return getters.GetInt(ma.obj.GetProperty("DirectionNumber"))
}

func (ma *RotatebleAdapater) SetDirection(n int) error {
	d, err := getters.GetInt(ma.obj.GetProperty("Direction"))
	if err != nil {
		return err
	}

	if n >= d {
		n = n - d
	}

	return ma.obj.SetProperty("DirectionNumber", n)
}
