package move

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw13/additions/vector"
	"otus-architecture/hw13/command/rotate"
	"otus-architecture/hw13/macrocommand"
	"otus-architecture/hw13/object"
	"testing"
)

func TestChangeVelocityOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))

	rotateAdpater := rotate.NewRotateAdapter(obj)
	rotateCommand := rotate.NewRotateCommand(rotateAdpater)

	moveAdapter := NewMoveAdapter(obj)
	moveCommand := NewMoveCommand(moveAdapter)

	macroCommand := macrocommand.NewMacroCommand(
		rotateCommand,
		moveCommand,
	)

	_, err := macroCommand.Execute()
	require.Equal(t, nil, err)

	velocityVector, err := obj.GetProperty("VelocityVector")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{-10, 0}, velocityVector)
}
