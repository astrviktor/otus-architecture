package hw04

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw04/command/checkfuel"
	"otus-architecture/hw04/object"
	"testing"
)

func TestCheckFuelOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Fuel", int(10))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))
	require.Equal(t, nil, err)

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)
	err = checkFuelCommand.Execute()

	require.Equal(t, nil, err)
}

func TestCheckFuelErrorFuelNotEnough(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Fuel", int(1))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)
	err = checkFuelCommand.Execute()

	require.Equal(t, checkfuel.ErrFuelNotEnough, err)
}

func TestCheckFuelErrorGetFuel(t *testing.T) {
	obj := object.New()

	checkFuelAdapter := checkfuel.NewCheckFuelAdapter(obj)
	checkFuelCommand := checkfuel.NewCheckFuelCommand(checkFuelAdapter)
	err := checkFuelCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}
