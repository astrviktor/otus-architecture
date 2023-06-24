package takeobject

import (
	"otus-architecture/hw11/additions/getters"
	"otus-architecture/hw11/object"
)

type TakeObjectAdapter struct {
	*object.Object
}

func NewTakeObjectAdapter(obj *object.Object) *TakeObjectAdapter {
	return &TakeObjectAdapter{obj}
}

func (a *TakeObjectAdapter) GetGameIdentifier() (int, error) {
	return getters.GetInt(a.GetProperty("GameIdentifier"))
}

func (a *TakeObjectAdapter) GetObjectIdentifier() (int, error) {
	return getters.GetInt(a.GetProperty("ObjectIdentifier"))
}
