package set

import (
	"otus-architecture/hw08/object"
)

type SetAdapter struct {
	*object.Object
}

func NewSetAdapter(obj *object.Object) *SetAdapter {
	return &SetAdapter{obj}
}
