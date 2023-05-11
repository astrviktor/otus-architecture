package rotate

import (
	"errors"
	"otus-architecture/hw03/additions/getters"
	"otus-architecture/hw03/object"
)

type RotateAdapterMock struct {
	*object.Object
}

func NewRotateAdapterMock(obj *object.Object) *RotateAdapterMock {
	return &RotateAdapterMock{obj}
}

func (a *RotateAdapterMock) GetDirection() (int, error) {
	return getters.GetInt(a.GetProperty("DirectionNumber"))
}

func (a *RotateAdapterMock) SetDirection(n int) error {
	return ErrDirectionFreeze
}

var (
	ErrDirectionFreeze = errors.New("direction freeze")
)
