package checkfuel

import (
	"otus-architecture/hw07/additions/getters"
	"otus-architecture/hw07/object"
)

type CheckFuelAdapter struct {
	*object.Object
}

func NewCheckFuelAdapter(obj *object.Object) *CheckFuelAdapter {
	return &CheckFuelAdapter{obj}
}

func (a *CheckFuelAdapter) GetFuel() (int, error) {
	return getters.GetInt(a.GetProperty("Fuel"))
}

func (a *CheckFuelAdapter) GetFuelUse() (int, error) {
	return getters.GetInt(a.GetProperty("FuelUse"))
}
