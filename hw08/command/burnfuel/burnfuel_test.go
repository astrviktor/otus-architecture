package burnfuel

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw07/object"
	"testing"
)

func TestBurnFuelOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Fuel", int(10))
	obj.SetProperty("FuelUse", int(3))

	burnFuelAdapter := NewBurnFuelAdapter(obj)
	burnFuelCommand := NewBurnFuelCommand(burnFuelAdapter)
	err := burnFuelCommand.Execute()

	require.Equal(t, nil, err)
}

func TestBurnFuelErrorFuelNotEnough(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Fuel", int(1))
	obj.SetProperty("FuelUse", int(3))

	burnFuelAdapter := NewBurnFuelAdapter(obj)
	burnFuelCommand := NewBurnFuelCommand(burnFuelAdapter)
	err := burnFuelCommand.Execute()

	require.Equal(t, ErrFuelNotEnough, err)
}

func TestBurnFuelErrorGetFuel(t *testing.T) {
	obj := object.New()

	burnFuelAdapter := NewBurnFuelAdapter(obj)
	burnFuelCommand := NewBurnFuelCommand(burnFuelAdapter)
	err := burnFuelCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}
