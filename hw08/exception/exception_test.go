package exception

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw08/additions/queue"
	"otus-architecture/hw08/additions/vector"
	"otus-architecture/hw08/command/move"
	"otus-architecture/hw08/object"
	"reflect"
	"testing"
)

func TestExceptionAllOk(t *testing.T) {
	expectedCommand := []string{
		"*move.MoveCommand",
	}

	stop := false
	exceptionHandler := NewExceptionRepeatLog()
	queue.CommandQueue = queue.New(100)

	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))

	moveCommand := move.NewMoveCommand(move.NewMoveAdapter(obj))
	queue.CommandQueue.Add(moveCommand)

	require.Equal(t, 1, queue.CommandQueue.GetLength())

	for !stop {
		if _, err := queue.CommandQueue.Peek(); err != nil {
			stop = true
			continue
		}

		cmd := queue.CommandQueue.Take()

		// testing
		require.Equal(t, expectedCommand[0], reflect.TypeOf(cmd).String())
		expectedCommand = expectedCommand[1:]

		_, err := cmd.Execute()
		if err != nil {
			exceptionHandler.Handle(cmd, err)
		}
	}

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{5, 12}, position)
}

func TestExceptionRunRepeatLog(t *testing.T) {
	expectedCommand := []string{
		"*move.MoveCommand",
		"*repeat.RepeatCommand",
		"*log.LogCommand",
	}

	stop := false
	exceptionHandler := NewExceptionRepeatLog()
	queue.CommandQueue = queue.New(100)

	obj := object.New()
	moveCommand := move.NewMoveCommand(move.NewMoveAdapter(obj))
	queue.CommandQueue.Add(moveCommand)

	require.Equal(t, 1, queue.CommandQueue.GetLength())

	for !stop {
		if _, err := queue.CommandQueue.Peek(); err != nil {
			stop = true
			continue
		}

		cmd := queue.CommandQueue.Take()

		// testing
		require.Equal(t, expectedCommand[0], reflect.TypeOf(cmd).String())
		expectedCommand = expectedCommand[1:]

		_, err := cmd.Execute()
		if err != nil {
			exceptionHandler.Handle(cmd, err)
		}
	}
}

func TestExceptionRunRepeatRepeatLog(t *testing.T) {
	expectedCommand := []string{
		"*move.MoveCommand",
		"*repeat.RepeatCommand",
		"*repeatcount.RepeatCountCommand",
		"*log.LogCommand",
	}

	stop := false
	exceptionHandler := NewExceptionRepeatRepeatLog()
	queue.CommandQueue = queue.New(100)

	obj := object.New()
	moveCommand := move.NewMoveCommand(move.NewMoveAdapter(obj))
	queue.CommandQueue.Add(moveCommand)

	require.Equal(t, 1, queue.CommandQueue.GetLength())

	for !stop {
		if _, err := queue.CommandQueue.Peek(); err != nil {
			stop = true
			continue
		}

		cmd := queue.CommandQueue.Take()

		// testing
		require.Equal(t, expectedCommand[0], reflect.TypeOf(cmd).String())
		expectedCommand = expectedCommand[1:]

		_, err := cmd.Execute()
		if err != nil {
			exceptionHandler.Handle(cmd, err)
		}
	}
}
