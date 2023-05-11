package log

import (
	"otus-architecture/hw07/additions/getters"
	"otus-architecture/hw07/object"
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
