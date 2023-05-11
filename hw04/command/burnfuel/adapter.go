package burnfuel

import (
	"otus-architecture/hw04/additions/getters"
	"otus-architecture/hw04/object"
)

type BurnFuelAdapter struct {
	*object.Object
}

func NewBurnFuelAdapter(obj *object.Object) *BurnFuelAdapter {
	return &BurnFuelAdapter{obj}
}

func (a *BurnFuelAdapter) GetFuel() (int, error) {
	return getters.GetInt(a.GetProperty("Fuel"))
}

func (a *BurnFuelAdapter) SetFuel(f int) error {
	return a.SetProperty("Fuel", f)
}

func (a *BurnFuelAdapter) GetFuelUse() (int, error) {
	return getters.GetInt(a.GetProperty("FuelUse"))
}
