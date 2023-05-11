package hw04

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw04/additions/vector"
	"otus-architecture/hw04/command/move"
	"otus-architecture/hw04/command/rotate"
	"otus-architecture/hw04/macrocommand"
	"otus-architecture/hw04/object"
	"testing"
)

func TestChangeVelocityOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	rotateAdpater := rotate.NewRotateAdapter(obj)
	rotateCommand := rotate.NewRotateCommand(rotateAdpater)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	macroCommand := macrocommand.NewMacroCommand(
		rotateCommand,
		moveCommand,
	)

	err = macroCommand.Execute()
	require.Equal(t, nil, err)

	velocityVector, err := obj.GetProperty("VelocityVector")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{-10, 0}, velocityVector)
}
