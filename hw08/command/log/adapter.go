package log

import (
	"otus-architecture/hw08/additions/getters"
	"otus-architecture/hw08/object"
)

type LogAdapter struct {
	*object.Object
}

func NewLogAdapter(obj *object.Object) *LogAdapter {
	return &LogAdapter{obj}
}

func (a *LogAdapter) GetError() (error, error) {
	return getters.GetError(a.GetProperty("Error"))
}
