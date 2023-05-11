package ioc

import (
	"otus-architecture/hw05/command"
	"otus-architecture/hw05/command/insertionsort"
	"otus-architecture/hw05/command/mergesort"
	"otus-architecture/hw05/command/selectionsort"
)

func SelectionSort(args ...interface{}) command.CommandInterface {
	data, ok := args[0].([]int)
	if !ok {
		return nil
	}

	return selectionsort.NewSelectionSortCommand(data)
}

func InsertionSort(args ...interface{}) command.CommandInterface {
	data, ok := args[0].([]int)
	if !ok {
		return nil
	}

	return insertionsort.NewInsertionSortCommand(data)
}

func MergeSort(args ...interface{}) command.CommandInterface {
	data, ok := args[0].([]int)
	if !ok {
		return nil
	}

	return mergesort.NewMergeSortCommand(data)
}
