package iocstrategies

import (
	"otus-architecture/hw08/additions/queue"
	"otus-architecture/hw08/additions/vector"
	"otus-architecture/hw08/command"
	"otus-architecture/hw08/command/makeobject"
	"otus-architecture/hw08/command/move"
	"otus-architecture/hw08/command/set"
	"otus-architecture/hw08/command/takeobject"
	"otus-architecture/hw08/object"
)

func Register(args ...interface{}) command.CommandInterface {
	key, ok := args[0].(string)
	if !ok {
		return nil
	}

	switch key {
	case "MakeObject":
		gameIdentifier, ok := args[1].(int)
		if !ok {
			return nil
		}

		objectIdentifier, ok := args[2].(int)
		if !ok {
			return nil
		}

		obj := object.New()
		obj.SetProperty("GameIdentifier", gameIdentifier)
		obj.SetProperty("ObjectIdentifier", objectIdentifier)
		makeObjectAdapter := makeobject.NewMakeObjectAdapter(obj)
		makeObjectCommand := makeobject.NewMakeObjectCommand(makeObjectAdapter)

		return makeObjectCommand
	}

	return nil
}

//obj, err := ioc.Resolve("TakeObject", message.GameID, message.ObjectID).Execute()

func TakeObject(args ...interface{}) command.CommandInterface {
	gameIdentifier, ok := args[0].(int)
	if !ok {
		return nil
	}

	objectIdentifier, ok := args[1].(int)
	if !ok {
		return nil
	}

	obj := object.New()
	obj.SetProperty("GameIdentifier", gameIdentifier)
	obj.SetProperty("ObjectIdentifier", objectIdentifier)

	takeObjectAdapter := takeobject.NewTakeObjectAdapter(obj)
	takeObjectCommand := takeobject.NewTakeObjectCommand(takeObjectAdapter)

	return takeObjectCommand
}

//_, err = ioc.Resolve("Position", obj, message.Args[0]).Execute()

func Position(args ...interface{}) command.CommandInterface {
	obj, ok := args[0].(*object.Object)
	if !ok {
		return nil
	}

	position, ok := args[1].(vector.Vector)
	if !ok {
		return nil
	}

	obj.SetProperty("Position", position)

	setAdapter := set.NewSetAdapter(obj)
	setCommand := set.NewSetCommand(setAdapter)

	return setCommand
}

// _, err = ioc.Resolve("Direction", obj, message.Args[1]).Execute()
func Direction(args ...interface{}) command.CommandInterface {
	obj, ok := args[0].(*object.Object)
	if !ok {
		return nil
	}

	position, ok := args[1].(int)
	if !ok {
		return nil
	}

	obj.SetProperty("Direction", position)

	setAdapter := set.NewSetAdapter(obj)
	setCommand := set.NewSetCommand(setAdapter)

	return setCommand
}

// _, err = ioc.Resolve("DirectionNumber", obj, message.Args[2]).Execute()
func DirectionNumber(args ...interface{}) command.CommandInterface {
	obj, ok := args[0].(*object.Object)
	if !ok {
		return nil
	}

	position, ok := args[1].(int)
	if !ok {
		return nil
	}

	obj.SetProperty("DirectionNumber", position)

	setAdapter := set.NewSetAdapter(obj)
	setCommand := set.NewSetCommand(setAdapter)

	return setCommand
}

// _, err = ioc.Resolve("Velocity", obj, message.Args[3]).Execute()
func Velocity(args ...interface{}) command.CommandInterface {
	obj, ok := args[0].(*object.Object)
	if !ok {
		return nil
	}

	position, ok := args[1].(int)
	if !ok {
		return nil
	}

	obj.SetProperty("Velocity", position)

	setAdapter := set.NewSetAdapter(obj)
	setCommand := set.NewSetCommand(setAdapter)

	return setCommand
}

// cmd := ioc.Resolve("MoveCommand", obj)
func MoveCommand(args ...interface{}) command.CommandInterface {
	obj, ok := args[0].(*object.Object)
	if !ok {
		return nil
	}

	moveAdapter := move.NewMoveAdapter(obj)
	moveCommand := move.NewMoveCommand(moveAdapter)

	return moveCommand
}

// _, err = ioc.Resolve("Queue", cmd).Execute()
func Queue(args ...interface{}) command.CommandInterface {
	cmd, ok := args[0].(command.CommandInterface)
	if !ok {
		return nil
	}

	queue.CommandQueue.Add(cmd)

	return &Nop{}
}
