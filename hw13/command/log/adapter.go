package log

import (
	"otus-architecture/hw13/additions/getters"
	"otus-architecture/hw13/object"
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
