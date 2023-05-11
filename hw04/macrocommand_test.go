package hw04

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw04/additions/vector"
	"otus-architecture/hw04/command/burnfuel"
	"otus-architecture/hw04/command/checkfuel"
	"otus-architecture/hw04/command/move"
	"otus-architecture/hw04/macrocommand"
	"otus-architecture/hw04/object"
	"testing"
)

func TestMacroCommendOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Fuel", int(10))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))
	require.Equal(t, nil, err)

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)

	macroCommand := macrocommand.NewMacroCommand(
		checkFuelCommand,
		moveCommand,
		burnFuelCommand,
	)

	err = macroCommand.Execute()
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

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Fuel", int(1))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))
	require.Equal(t, nil, err)

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)

	macroCommand := macrocommand.NewMacroCommand(
		checkFuelCommand,
		moveCommand,
		burnFuelCommand,
	)

	err = macroCommand.Execute()
	require.Equal(t, checkfuel.ErrFuelNotEnough, err)

}
