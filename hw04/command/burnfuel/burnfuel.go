package burnfuel

import "errors"

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

func (c *BurnFuelCommand) Execute() error {
	var fuel, fuelUse int
	var err error

	if fuel, err = c.obj.GetFuel(); err != nil {
		return err
	}

	if fuelUse, err = c.obj.GetFuelUse(); err != nil {
		return err
	}

	if fuel < fuelUse {
		return ErrFuelNotEnough
	}

	fuel = fuel - fuelUse

	if err = c.obj.SetFuel(fuel); err != nil {
		return err
	}

	return nil
}

var (
	ErrFuelNotEnough = errors.New("fuel not enough")
)
