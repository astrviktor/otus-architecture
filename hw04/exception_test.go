package hw04

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw04/additions/queue"
	"otus-architecture/hw04/additions/vector"
	"otus-architecture/hw04/command/move"
	"otus-architecture/hw04/exception"
	"otus-architecture/hw04/object"
	"reflect"
	"testing"
)

func TestExceptionAllOk(t *testing.T) {
	expectedCommand := []string{
		"*move.MoveCommand",
	}

	stop := false
	exceptionHandler := exception.NewExceptionRepeatLog()
	queue.CommandQueue = queue.New(100)

	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

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

		err := cmd.Execute()
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
	exceptionHandler := exception.NewExceptionRepeatLog()
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

		err := cmd.Execute()
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
	exceptionHandler := exception.NewExceptionRepeatRepeatLog()
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

		err := cmd.Execute()
		if err != nil {
			exceptionHandler.Handle(cmd, err)
		}
	}
}
