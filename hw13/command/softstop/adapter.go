package softstop

import (
	"otus-architecture/hw13/additions/getters"
	"otus-architecture/hw13/object"
)

type SoftStopAdapter struct {
	*object.Object
}

func NewSoftStopAdapter(obj *object.Object) *SoftStopAdapter {
	return &SoftStopAdapter{obj}
}

func (a *SoftStopAdapter) GetSoftStopChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("SoftStopChannel"))
}
