package set

import (
	"otus-architecture/hw11/object"
)

type SetAdapter struct {
	*object.Object
}

func NewSetAdapter(obj *object.Object) *SetAdapter {
	return &SetAdapter{obj}
}
