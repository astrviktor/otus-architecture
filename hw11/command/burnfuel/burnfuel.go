package burnfuel

import (
	"errors"
	"otus-architecture/hw11/object"
)

type BurnFuelInterface interface {
	GetFuel() (int, error)
	SetFuel(f int) error
	GetFuelUse() (int, error)
}

type BurnFuelCommand struct {
	obj BurnFuelInterface
}

func NewBurnFuelCommand(obj BurnFuelInterface) *BurnFuelCommand {
	return &BurnFuelCommand{obj: obj}
}

func (c *BurnFuelCommand) Execute() (*object.Object, error) {
	var fuel, fuelUse int
	var err error

	if fuel, err = c.obj.GetFuel(); err != nil {
		return nil, err
	}

	if fuelUse, err = c.obj.GetFuelUse(); err != nil {
		return nil, err
	}

	if fuel < fuelUse {
		return nil, ErrFuelNotEnough
	}

	fuel = fuel - fuelUse

	if err = c.obj.SetFuel(fuel); err != nil {
		return nil, err
	}

	return nil, nil
}

var (
	ErrFuelNotEnough = errors.New("fuel not enough")
)
