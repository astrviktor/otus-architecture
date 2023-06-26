package checkfuel

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw12/object"
	"testing"
)

func TestCheckFuelOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Fuel", int(10))
	obj.SetProperty("FuelUse", int(3))

	checkFuelAdapter := NewCheckFuelAdapter(obj)
	checkFuelCommand := NewCheckFuelCommand(checkFuelAdapter)
	err := checkFuelCommand.Execute()

	require.Equal(t, nil, err)
}

func TestCheckFuelErrorFuelNotEnough(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Fuel", int(1))
	obj.SetProperty("FuelUse", int(3))

	checkFuelAdapter := NewCheckFuelAdapter(obj)
	checkFuelCommand := NewCheckFuelCommand(checkFuelAdapter)
	err := checkFuelCommand.Execute()

	require.Equal(t, ErrFuelNotEnough, err)
}

func TestCheckFuelErrorGetFuel(t *testing.T) {
	obj := object.New()

	checkFuelAdapter := NewCheckFuelAdapter(obj)
	checkFuelCommand := NewCheckFuelCommand(checkFuelAdapter)
	err := checkFuelCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}
