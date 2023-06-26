package move

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/object"
	"testing"
)

func TestMoveOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))

	moveAdapter := NewMoveAdapter(obj)
	moveCommand := NewMoveCommand(moveAdapter)
	_, err := moveCommand.Execute()

	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{5, 12}, position)
}

func TestMoveErrorGetPosition(t *testing.T) {
	obj := object.New()

	moveAdapter := NewMoveAdapter(obj)
	moveCommand := NewMoveCommand(moveAdapter)
	_, err := moveCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

func TestMoveErrorGetVelocity(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))

	moveAdapter := NewMoveAdapter(obj)
	moveCommand := NewMoveCommand(moveAdapter)
	_, err := moveCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

//func TestMoveErrorFreeze(t *testing.T) {
//	obj := object.New()
//
//	obj.SetProperty("Position", vector.Vector{12, 5})
//	obj.SetProperty("Direction", int(8))
//	obj.SetProperty("DirectionNumber", int(3))
//	obj.SetProperty("Velocity", int(10))
//
//	moveAdapter := NewMoveAdapterMock(obj)
//	moveCommand := NewMoveCommand(moveAdapter)
//	_, err := moveCommand.Execute()
//
//	require.Equal(t, ErrPositionFreeze, err)
//}
