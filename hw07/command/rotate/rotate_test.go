package rotate

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw07/additions/vector"
	"otus-architecture/hw07/object"
	"testing"
)

func TestRotateOk(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))

	rotateAdpater := NewRotateAdapter(obj)
	rotateCommand := NewRotateCommand(rotateAdpater)
	err := rotateCommand.Execute()

	require.Equal(t, nil, err)

	n, err := obj.GetProperty("DirectionNumber")

	require.Equal(t, nil, err)
	require.Equal(t, 4, n)
}

func TestRotateErrorGetDirectionNumber(t *testing.T) {
	obj := object.New()

	rotateAdpater := NewRotateAdapter(obj)
	rotateCommand := NewRotateCommand(rotateAdpater)
	err := rotateCommand.Execute()

	require.Equal(t, object.ErrNoDataByKey, err)
}

func TestRotateErrorFreeze(t *testing.T) {
	obj := object.New()

	obj.SetProperty("Position", vector.Vector{12, 5})
	obj.SetProperty("Direction", int(8))
	obj.SetProperty("DirectionNumber", int(3))
	obj.SetProperty("Velocity", int(10))

	rotateAdpater := NewRotateAdapterMock(obj)
	rotateCommand := NewRotateCommand(rotateAdpater)
	err := rotateCommand.Execute()

	require.Equal(t, ErrDirectionFreeze, err)
}
