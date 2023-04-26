package rotate

import (
	"otus-architecture/hw04/additions/getters"
	"otus-architecture/hw04/object"
)

type RotateAdapter struct {
	*object.Object
}

func NewRotateAdapter(obj *object.Object) *RotateAdapter {
	return &RotateAdapter{obj}
}

func (a *RotateAdapter) GetDirection() (int, error) {
	return getters.GetInt(a.GetProperty("DirectionNumber"))
}

func (a *RotateAdapter) SetDirection(n int) error {
	d, err := getters.GetInt(a.GetProperty("Direction"))
	if err != nil {
		return err
	}

	if n >= d {
		n = n - d
	}

	return a.SetProperty("DirectionNumber", n)
}
