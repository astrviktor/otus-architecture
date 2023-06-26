package hardstop

import (
	"otus-architecture/hw11/additions/getters"
	"otus-architecture/hw11/object"
)

type HardStopAdapter struct {
	*object.Object
}

func NewHardStopAdapter(obj *object.Object) *HardStopAdapter {
	return &HardStopAdapter{obj}
}

func (a *HardStopAdapter) GetHardStopChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("HardStopChannel"))
}
