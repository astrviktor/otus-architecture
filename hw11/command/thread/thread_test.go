package thread

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw11/additions/queue"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/command/hardstop"
	"otus-architecture/hw11/command/move"
	"otus-architecture/hw11/command/softstop"
	"otus-architecture/hw11/object"
	"testing"
	"time"
)

func TestThreadRunMove3TimesOk(t *testing.T) {
	// move command
	moveObject := object.New()

	moveObject.SetProperty("Position", vector.Vector{0, 0})
	moveObject.SetProperty("Direction", int(8))
	moveObject.SetProperty("DirectionNumber", int(0))
	moveObject.SetProperty("Velocity", int(1))

	moveAdapter := move.NewMoveAdapter(moveObject)
	moveCommand := move.NewMoveCommand(moveAdapter)

	// thread command
	threadObject := object.New()

	queue := make(chan queue.Element, 10)
	hardStopChannel := make(chan struct{}, 1)
	softStopChannel := make(chan struct{}, 1)
	endChannel := make(chan struct{})

	threadObject.SetProperty("Queue", queue)
	threadObject.SetProperty("HardStopChannel", hardStopChannel)
	threadObject.SetProperty("SoftStopChannel", softStopChannel)
	threadObject.SetProperty("EndChannel", endChannel)

	threadAdapter := NewThreadAdapter(threadObject)
	threadCommand := NewThreadCommand(threadAdapter)

	_, err := threadCommand.Execute()
	require.Equal(t, nil, err)

	// to queue added: 3 move command
	queue <- moveCommand
	queue <- moveCommand
	queue <- moveCommand

	// thread wait commands, test waiting 5 seconds and end
	time.Sleep(5 * time.Second)

	// checked "Position" result of execute 3 move command
	position, err := moveObject.GetProperty("Position")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{3, 0}, position)
}

func TestThreadRunMove3TimesHardStop(t *testing.T) {
	// move command
	moveObject := object.New()

	moveObject.SetProperty("Position", vector.Vector{0, 0})
	moveObject.SetProperty("Direction", int(8))
	moveObject.SetProperty("DirectionNumber", int(0))
	moveObject.SetProperty("Velocity", int(1))

	moveAdapter := move.NewMoveAdapter(moveObject)
	moveCommand := move.NewMoveCommand(moveAdapter)

	// thread command
	threadObject := object.New()

	queue := make(chan queue.Element, 10)
	hardStopChannel := make(chan struct{}, 1)
	softStopChannel := make(chan struct{}, 1)
	endChannel := make(chan struct{})

	threadObject.SetProperty("Queue", queue)
	threadObject.SetProperty("HardStopChannel", hardStopChannel)
	threadObject.SetProperty("SoftStopChannel", softStopChannel)
	threadObject.SetProperty("EndChannel", endChannel)

	threadAdapter := NewThreadAdapter(threadObject)
	threadCommand := NewThreadCommand(threadAdapter)

	_, err := threadCommand.Execute()
	require.Equal(t, nil, err)

	// hard stop command
	hardStopObject := object.New()
	hardStopObject.SetProperty("HardStopChannel", hardStopChannel)

	hardStopAdapter := hardstop.NewHardStopAdapter(hardStopObject)
	hardStopCommand := hardstop.NewHardStopCommand(hardStopAdapter)

	// to queue added: 2 move command, hard stop command, move command
	queue <- moveCommand
	queue <- moveCommand
	queue <- hardStopCommand
	queue <- moveCommand

	// sync thread stop
	<-endChannel

	// checked "Position" result of execute only 2 move command
	position, err := moveObject.GetProperty("Position")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{2, 0}, position)
}

func TestThreadRunMove3TimesSoftStop(t *testing.T) {
	// move command
	moveObject := object.New()

	moveObject.SetProperty("Position", vector.Vector{0, 0})
	moveObject.SetProperty("Direction", int(8))
	moveObject.SetProperty("DirectionNumber", int(0))
	moveObject.SetProperty("Velocity", int(1))

	moveAdapter := move.NewMoveAdapter(moveObject)
	moveCommand := move.NewMoveCommand(moveAdapter)

	// thread command
	threadObject := object.New()

	queue := make(chan queue.Element, 10)
	hardStopChannel := make(chan struct{}, 1)
	softStopChannel := make(chan struct{}, 1)
	endChannel := make(chan struct{})

	threadObject.SetProperty("Queue", queue)
	threadObject.SetProperty("HardStopChannel", hardStopChannel)
	threadObject.SetProperty("SoftStopChannel", softStopChannel)
	threadObject.SetProperty("EndChannel", endChannel)

	threadAdapter := NewThreadAdapter(threadObject)
	threadCommand := NewThreadCommand(threadAdapter)

	_, err := threadCommand.Execute()
	require.Equal(t, nil, err)

	// soft stop command
	softStopObject := object.New()
	softStopObject.SetProperty("SoftStopChannel", softStopChannel)

	softStopAdapter := softstop.NewSoftStopAdapter(softStopObject)
	softStopCommand := softstop.NewSoftStopCommand(softStopAdapter)

	// to queue added: 2 move command, soft stop command, move command
	queue <- moveCommand
	queue <- moveCommand
	queue <- softStopCommand
	queue <- moveCommand

	// sync thread stop
	<-endChannel

	// checked "Position" result of execute all 3 move command
	position, err := moveObject.GetProperty("Position")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{3, 0}, position)
}
