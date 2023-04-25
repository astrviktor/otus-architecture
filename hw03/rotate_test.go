package hw03

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw03/additions/vector"
	"otus-architecture/hw03/command/rotate"
	"otus-architecture/hw03/object"
	"testing"
)

func TestRotateOk(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	rotateAdpater := rotate.NewRotateAdapter(obj)
	rotateCommand := rotate.NewRotateCommand(rotateAdpater)
	err = rotateCommand.Execute()

	require.Equal(t, nil, err)

	n, err := obj.GetProperty("DirectionNumber")

	require.Equal(t, nil, err)
	require.Equal(t, 4, n)
}

func TestRotateErrorGetDirectionNumber(t *testing.T) {
	obj := object.New()

	rotateAdpater := rotate.NewRotateAdapter(obj)
	rotateCommand := rotate.NewRotateCommand(rotateAdpater)
	err := rotateCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

func TestRotateErrorFreeze(t *testing.T) {
	obj := object.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	rotateAdpater := rotate.NewRotateAdapterMock(obj)
	rotateCommand := rotate.NewRotateCommand(rotateAdpater)
	err = rotateCommand.Execute()

	require.Equal(t, rotate.ErrDirectionFreeze, err)
}
