package ioc

import (
	"otus-architecture/hw11/command"
	"otus-architecture/hw11/ioc/iocstrategies"
)

type StrategyFunc func(args ...interface{}) command.CommandInterface

type IoC struct {
	strategy map[string]StrategyFunc
}

func New() *IoC {
	strategy := make(map[string]StrategyFunc)

	strategy["Register"] = iocstrategies.Register
	strategy["TakeObject"] = iocstrategies.TakeObject
	strategy["Position"] = iocstrategies.Position
	strategy["Direction"] = iocstrategies.Direction
	strategy["DirectionNumber"] = iocstrategies.DirectionNumber
	strategy["Velocity"] = iocstrategies.Velocity
	strategy["MoveCommand"] = iocstrategies.MoveCommand
	strategy["Queue"] = iocstrategies.Queue

	return &IoC{strategy: strategy}
}

func (ioc *IoC) Resolve(key string, args ...interface{}) command.CommandInterface {
	strategy, ok := ioc.strategy[key]
	if !ok {
		return nil
	}

	cmd := strategy(args...)
	return cmd
}
