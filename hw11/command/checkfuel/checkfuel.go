package checkfuel

import (
	"errors"
	"otus-architecture/hw11/object"
)

type CheckFuelInterface interface {
	GetFuel() (int, error)
	GetFuelUse() (int, error)
}

type CheckFuelCommand struct {
	obj CheckFuelInterface
}

func NewCheckFuelCommand(obj CheckFuelInterface) *CheckFuelCommand {
	return &CheckFuelCommand{obj: obj}
}

func (c *CheckFuelCommand) Execute() (*object.Object, error) {
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

	return nil, nil
}

var (
	ErrFuelNotEnough = errors.New("fuel not enough")
)
