package run

import (
	"otus-architecture/hw12/additions/getters"
	"otus-architecture/hw12/object"
)

type RunAdapter struct {
	*object.Object
}

func NewRunAdapter(obj *object.Object) *RunAdapter {
	return &RunAdapter{obj}
}

func (a *RunAdapter) GetRunChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("RunChannel"))
}
