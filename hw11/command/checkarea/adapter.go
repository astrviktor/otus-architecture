package checkarea

import (
	"math"
	"otus-architecture/hw11/additions/getters"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/command"
	"otus-architecture/hw11/object"
)

type CheckAreaAdapter struct {
	*object.Object
}

func NewCheckAreaAdapter(obj *object.Object) *CheckAreaAdapter {
	return &CheckAreaAdapter{obj}
}

func (a *CheckAreaAdapter) GetAreasNames() ([]string, error) {
	return getters.GetStringSlice(a.GetProperty("AreasNames"))
}

func (a *CheckAreaAdapter) GetAreaForName(name string) (vector.Vector, error) {
	return getters.GetVector(a.GetProperty("Position"))
}

type Areafunc = func(position vector.Vector) vector.Vector

func (a *CheckAreaAdapter) GetAreaForNameFromPosition(name string) (vector.Vector, error) {
	return getters.GetVector(a.GetProperty("Position"))
}

func (a *CheckAreaAdapter) SetAreaForName(name string, area vector.Vector) error {
	return nil
}

func (a *CheckAreaAdapter) MakeCollisionMacroCommand(name string, area vector.Vector) (command.CommandInterface, error) {
	return nil, nil
}

func (a *CheckAreaAdapter) SetCollisionMacroCommand(name string, area vector.Vector, cmd command.CommandInterface) error {
	return nil
}

func (a *CheckAreaAdapter) GetPosition() (vector.Vector, error) {
	return getters.GetVector(a.GetProperty("Position"))
}

type objects = []object.Object
type area = map[vector.Vector]objects

var areas = make(map[string]area)

func isCollision(p1 vector.Vector, r1 float64, p2 vector.Vector, r2 float64) bool {
	r := math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))

	return r < r1+r2
}
