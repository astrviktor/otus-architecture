package thread

import (
	"otus-architecture/hw12/additions/getters"
	"otus-architecture/hw12/additions/queue"
	"otus-architecture/hw12/object"
)

type ThreadAdapter struct {
	*object.Object
}

func NewThreadAdapter(obj *object.Object) *ThreadAdapter {
	return &ThreadAdapter{obj}
}

func (a *ThreadAdapter) GetHardStopChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("HardStopChannel"))
}

func (a *ThreadAdapter) GetSoftStopChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("SoftStopChannel"))
}

func (a *ThreadAdapter) GetEndChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("EndChannel"))
}

func (a *ThreadAdapter) GetQueueChannel() (chan queue.Element, error) {
	return getters.GetQueueChannel(a.GetProperty("Queue"))
}
