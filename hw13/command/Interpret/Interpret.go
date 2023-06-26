package interpret

import (
	"otus-architecture/hw13/broker"
	"otus-architecture/hw13/ioc"
)

type InterpretInterface interface {
	GetIoC() (ioc.IoC, error)
	GetMessage() (broker.Message, error)
}

type InterpretCommand struct {
	obj *InterpretAdapter
}

func NewInterpretCommand(obj *InterpretAdapter) *InterpretCommand {
	return &InterpretCommand{obj: obj}
}

func (c *InterpretCommand) Execute() error {
	ioc, err := c.obj.GetIoC()
	if err != nil {
		return err
	}

	message, err := c.obj.GetMessage()
	if err != nil {
		return err
	}

	//var obj = IoC.Resolve("Игровые объекты", "548") // "548" получено из входящего сообщения
	obj, err := ioc.Resolve("TakeObject", message.GameID, message.ObjectID).Execute()
	if err != nil {
		return err
	}

	commandName := CommandResolve(message.OperationID)
	if commandName == "MoveCommand" {
		//IoC.Resolve("Установить начальное значение скорости", obj, 2) // значение 2 получено из args переданного в сообщении
		_, err = ioc.Resolve("Position", obj, message.Args[0]).Execute()
		_, err = ioc.Resolve("Direction", obj, message.Args[1]).Execute()
		_, err = ioc.Resolve("DirectionNumber", obj, message.Args[2]).Execute()
		_, err = ioc.Resolve("Velocity", obj, message.Args[3]).Execute()

		//var cmd = IoC.Resolve("Движение по прямой", obj) // Решение, что нужно выполнить движение по прямой получено из сообщения
		cmd := ioc.Resolve("MoveCommand", obj)

		//IoC.Resolve("Очередь команд", cmd).Execute() // Выполняем команду, которая поместит команду cmd в очередь команд игры.
		_, err = ioc.Resolve("Queue", cmd).Execute()
	}

	return nil
}

func CommandResolve(id int) string {
	commands := make(map[int]string)
	commands[1] = "MoveCommand"

	command, ok := commands[id]
	if !ok {
		return "None"
	}

	return command
}
