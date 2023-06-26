package macrocommand

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw13/additions/vector"
	"otus-architecture/hw13/command/burnfuel"
	"otus-architecture/hw13/command/checkfuel"
	"otus-architecture/hw13/command/move"
	"otus-architecture/hw13/object"
	"testing"
)

func TestMacroCommendOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))
	obj.SetProperty("Fuel", int(10))
	obj.SetProperty("FuelUse", int(3))

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)

	macroCommand := NewMacroCommand(
		checkFuelCommand,
		moveCommand,
		burnFuelCommand,
	)

	_, err := macroCommand.Execute()
	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")
	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{5, 12}, position)

	fuel, err := obj.GetProperty("Fuel")
	require.Equal(t, nil, err)
	require.Equal(t, int(7), fuel)
}

func TestMacroCommendErrorFuelNotEnough(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))
	obj.SetProperty("Fuel", int(1))
	obj.SetProperty("FuelUse", int(3))

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)

	macroCommand := NewMacroCommand(
		checkFuelCommand,
		moveCommand,
		burnFuelCommand,
	)

	_, err := macroCommand.Execute()
	require.Equal(t, checkfuel.ErrFuelNotEnough, err)

}
