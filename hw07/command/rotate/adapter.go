package rotate

import (
	"otus-architecture/hw07/additions/getters"
	"otus-architecture/hw07/object"
)

type RotateAdapter struct {
	obj *object.Object
}

func NewRotateAdapter(obj *object.Object) *RotateAdapter {
	return &RotateAdapter{obj: obj}
}

func (a *RotateAdapter) GetObject() *object.Object {
	return a.obj
}

func (a *RotateAdapter) GetDirection() (int, error) {
	return getters.GetInt(a.obj.GetProperty("DirectionNumber"))
}

func (a *RotateAdapter) SetDirection(n int) error {
	d, err := getters.GetInt(a.obj.GetProperty("Direction"))
	if err != nil {
		return err
	}

	if n >= d {
		n = n - d
	}

	a.obj.SetProperty("DirectionNumber", n)

	return nil
}
