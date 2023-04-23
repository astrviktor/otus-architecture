package hw02

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw02/rotateble_adapter"
	rotateble_adapater_mock "otus-architecture/hw02/rotateble_adapter/rotateble_adapter_mock"
	"otus-architecture/hw02/uobject"
	"otus-architecture/hw02/vector"
	"testing"
)

func TestRotateOk(t *testing.T) {
	obj := uobject.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	rotatebleAdpater := rotateble_adapater.New(obj)
	rotate := Rotate{}
	err = rotate.Execute(&rotatebleAdpater)

	require.Equal(t, nil, err)

	n, err := obj.GetProperty("DirectionNumber")

	require.Equal(t, nil, err)
	require.Equal(t, 4, n)
}

func TestRotateErrorGetDirectionNumber(t *testing.T) {
	obj := uobject.New()

	rotatebleAdpater := rotateble_adapater.New(obj)
	rotate := Rotate{}
	err := rotate.Execute(&rotatebleAdpater)

	require.Equal(t, uobject.ErrNoDataByKey, err)
}

func TestRotateFreeze(t *testing.T) {
	obj := uobject.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	rotatebleAdpater := rotateble_adapater_mock.New(obj)
	rotate := Rotate{}
	err = rotate.Execute(&rotatebleAdpater)

	require.Equal(t, nil, err)

	n, err := obj.GetProperty("DirectionNumber")

	require.Equal(t, nil, err)
	require.Equal(t, 3, n)
}
