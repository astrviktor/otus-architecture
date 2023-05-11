package hw03

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw03/additions/vector"
	"otus-architecture/hw03/command/move"
	"otus-architecture/hw03/object"
	"testing"
)

func TestMoveOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)
	err = moveCommand.Execute()

	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{5, 12}, position)
}

func TestMoveErrorGetPosition(t *testing.T) {
	obj := object.New()

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)
	err := moveCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

func TestMoveErrorGetVelocity(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)
	err = moveCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

func TestMoveErrorFreeze(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	moveAdapter := move.NewMoveAdapterMock(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)
	err = moveCommand.Execute()

	require.Equal(t, move.ErrPositionFreeze, err)
}
