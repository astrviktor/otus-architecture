package hw04

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw04/command/burnfuel"
	"otus-architecture/hw04/object"
	"testing"
)

func TestBurnFuelOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Fuel", int(10))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))
	require.Equal(t, nil, err)

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)
	err = burnFuelCommand.Execute()

	require.Equal(t, nil, err)
}

func TestBurnFuelErrorFuelNotEnough(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Fuel", int(1))
	require.Equal(t, nil, err)
	err = obj.SetProperty("FuelUse", int(3))

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)
	err = burnFuelCommand.Execute()

	require.Equal(t, burnfuel.ErrFuelNotEnough, err)
}

func TestBurnFuelErrorGetFuel(t *testing.T) {
	obj := object.New()

	burnFuelAdapter := burnfuel.NewBurnFuelAdapter(obj)
	burnFuelCommand := burnfuel.NewBurnFuelCommand(burnFuelAdapter)
	err := burnFuelCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}
