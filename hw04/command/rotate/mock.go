package rotate

import (
	"errors"
	"otus-architecture/hw04/additions/getters"
	"otus-architecture/hw04/object"
)

type RotateAdapterMock struct {
	obj *object.Object
}

func NewRotateAdapterMock(obj *object.Object) *RotateAdapterMock {
	return &RotateAdapterMock{obj}
}

func (a *RotateAdapterMock) GetObject() *object.Object {
	return a.obj
}

func (a *RotateAdapterMock) GetDirection() (int, error) {
	return getters.GetInt(a.obj.GetProperty("DirectionNumber"))
}

func (a *RotateAdapterMock) SetDirection(n int) error {
	return ErrDirectionFreeze
}

var (
	ErrDirectionFreeze = errors.New("direction freeze")
)
