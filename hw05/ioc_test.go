package main

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw05/ioc"
	"reflect"
	"testing"
)

func TestSelectionSortOk(t *testing.T) {
	data := []int{2, 3, 1, 8, 4}
	expected := []int{1, 2, 3, 4, 8}

	factory := ioc.New()
	command := factory.Resolve("SelectionSort", data)

	require.NotEqual(t, nil, command)
	require.Equal(t, "*selectionsort.SelectionSortCommand",
		reflect.TypeOf(command).String())

	if command != nil {
		err := command.Execute()
		require.Equal(t, nil, err)
	}

	require.Equal(t, expected, data)

}

func TestInsertionSortOk(t *testing.T) {
	data := []int{2, 3, 1, 8, 4}
	expected := []int{1, 2, 3, 4, 8}

	factory := ioc.New()
	command := factory.Resolve("InsertionSort", data)

	require.NotEqual(t, nil, command)
	require.Equal(t, "*insertionsort.InsertionSortCommand",
		reflect.TypeOf(command).String())

	if command != nil {
		err := command.Execute()
		require.Equal(t, nil, err)
	}

	require.Equal(t, expected, data)
}

func TestMergeSortOk(t *testing.T) {
	data := []int{2, 3, 1, 8, 4}
	expected := []int{1, 2, 3, 4, 8}

	factory := ioc.New()
	command := factory.Resolve("MergeSort", data)

	require.NotEqual(t, nil, command)
	require.Equal(t, "*mergesort.MergeSortCommand",
		reflect.TypeOf(command).String())

	if command != nil {
		err := command.Execute()
		require.Equal(t, nil, err)
	}

	require.Equal(t, expected, data)
}
