package rotateble_adapater_mock

import (
	"otus-architecture/hw02/getters"
	"otus-architecture/hw02/uobject"
)

type RotatebleAdapaterMock struct {
	obj *uobject.UObject
}

func New(obj *uobject.UObject) RotatebleAdapaterMock {
	return RotatebleAdapaterMock{obj: obj}
}

func (ma *RotatebleAdapaterMock) GetDirection() (int, error) {
	return getters.GetInt(ma.obj.GetProperty("DirectionNumber"))
}

func (ma *RotatebleAdapaterMock) SetDirection(n int) error {
	return nil
}
