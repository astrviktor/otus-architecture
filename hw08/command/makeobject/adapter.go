package makeobject

import (
	"otus-architecture/hw08/additions/getters"
	"otus-architecture/hw08/object"
)

type MakeObjectAdapter struct {
	*object.Object
}

func NewMakeObjectAdapter(obj *object.Object) *MakeObjectAdapter {
	return &MakeObjectAdapter{obj}
}

func (a *MakeObjectAdapter) GetGameIdentifier() (int, error) {
	return getters.GetInt(a.GetProperty("GameIdentifier"))
}

func (a *MakeObjectAdapter) GetObjectIdentifier() (int, error) {
	return getters.GetInt(a.GetProperty("ObjectIdentifier"))
}
