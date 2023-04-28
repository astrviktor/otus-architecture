package ioc

import (
	"otus-architecture/hw05/command"
)

type StrategyFunc func(args ...interface{}) command.CommandInterface

type IoC struct {
	strategy map[string]StrategyFunc
}

func New() *IoC {
	strategy := make(map[string]StrategyFunc)

	strategy["SelectionSort"] = SelectionSort
	strategy["InsertionSort"] = InsertionSort
	strategy["MergeSort"] = MergeSort

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
