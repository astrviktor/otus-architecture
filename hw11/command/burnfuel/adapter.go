package burnfuel

import (
	"otus-architecture/hw11/additions/getters"
	"otus-architecture/hw11/object"
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
	a.SetProperty("Fuel", f)
	return nil
}

func (a *BurnFuelAdapter) GetFuelUse() (int, error) {
	return getters.GetInt(a.GetProperty("FuelUse"))
}
