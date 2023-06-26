package moveto

import (
	"otus-architecture/hw12/additions/getters"
	"otus-architecture/hw12/additions/queue"
	"otus-architecture/hw12/object"
)

type MoveToAdapter struct {
	*object.Object
}

func NewMoveToAdapter(obj *object.Object) *MoveToAdapter {
	return &MoveToAdapter{obj}
}

func (a *MoveToAdapter) GetMoveToQueue() (chan queue.Element, error) {
	return getters.GetQueueChannel(a.GetProperty("MoveToQueue"))
}

func (a *MoveToAdapter) GetQueue() (chan queue.Element, error) {
	return getters.GetQueueChannel(a.GetProperty("Queue"))
}

func (a *MoveToAdapter) GetRunChannel() (chan struct{}, error) {
	return getters.GetStructChannel(a.GetProperty("RunChannel"))
}
