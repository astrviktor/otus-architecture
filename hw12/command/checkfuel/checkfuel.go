package checkfuel

import "errors"

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

func (c *CheckFuelCommand) Execute() error {
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

	return nil
}

var (
	ErrFuelNotEnough = errors.New("fuel not enough")
)
