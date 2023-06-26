package checkarea

import (
	"otus-architecture/hw11/object"
)

type CheckAreaCommand struct {
	obj *CheckAreaAdapter
}

func NewCheckAreaCommand(obj *CheckAreaAdapter) *CheckAreaCommand {
	return &CheckAreaCommand{obj: obj}
}

func (c *CheckAreaCommand) Execute() (*object.Object, error) {
	//получить список игровых полей для объекта
	names, err := c.obj.GetAreasNames()
	if err != nil {
		return nil, err
	}

	//для всех игровых полей
	for idx := range names {
		// получить текущую окресность объекта
		area, err := c.obj.GetAreaForName(names[idx])
		if err != nil {
			return nil, err
		}

		// посчитать новую окресность объекта по координатам и названию поля
		areaNew, err := c.obj.GetAreaForNameFromPosition(names[idx])
		if err != nil {
			return nil, err
		}

		// сравнить окресности
		if area != areaNew {
			// если окресности разные

			// сделать новую макрокоманду для проверуки коллизий
			cmd, err := c.obj.MakeCollisionMacroCommand(names[idx], areaNew)
			if err != nil {
				return nil, err
			}

			// поместить новую макрокоманду взамен старой
			err = c.obj.SetCollisionMacroCommand(names[idx], areaNew, cmd)
			if err != nil {
				return nil, err
			}

			// удалить объект из старой окресности и добавить в новую
			err = c.obj.SetAreaForName(names[idx], areaNew)
			if err != nil {
				return nil, err
			}
		}
	}

	return c.obj.Object, nil
}
