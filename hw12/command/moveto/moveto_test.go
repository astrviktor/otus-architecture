package moveto

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw12/additions/queue"
	"otus-architecture/hw12/additions/vector"
	"otus-architecture/hw12/command/move"
	"otus-architecture/hw12/object"
	"testing"
)

func TestMoveToOk(t *testing.T) {
	obj := object.New()

	Queue := make(chan queue.Element, 10)
	MoveToQueue := make(chan queue.Element, 10)
	RunChannel := make(chan struct{}, 1)

	obj.SetProperty("Queue", Queue)
	obj.SetProperty("MoveToQueue", MoveToQueue)
	obj.SetProperty("RunChannel", RunChannel)

	testStopChannel := make(chan struct{}, 1)

	go func() {
		<-MoveToQueue
		close(testStopChannel)
	}()

	moveToAdapter := NewMoveToAdapter(obj)
	moveToCommand := NewMoveToCommand(moveToAdapter)

	err := moveToCommand.Execute()
	require.Equal(t, nil, err)

	// move command
	moveObject := object.New()

	moveObject.SetProperty("Position", vector.Vector{0, 0})
	moveObject.SetProperty("Direction", int(8))
	moveObject.SetProperty("DirectionNumber", int(0))
	moveObject.SetProperty("Velocity", int(1))

	moveAdapter := move.NewMoveAdapter(moveObject)
	moveCommand := move.NewMoveCommand(moveAdapter)

	Queue <- moveCommand

	<-testStopChannel
}
