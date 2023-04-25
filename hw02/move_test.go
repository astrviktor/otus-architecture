package hw02

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw02/movable_adapater"
	"otus-architecture/hw02/movable_adapater/movable_adapater_mock"
	"otus-architecture/hw02/uobject"
	"otus-architecture/hw02/vector"
	"testing"
)

func TestMoveOk(t *testing.T) {
	obj := uobject.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	movableAdpater := movable_adapater.New(obj)
	move := Move{}
	err = move.Execute(&movableAdpater)

	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{5, 12}, position)
}

func TestMoveErrorGetPosition(t *testing.T) {
	obj := uobject.New()

	movableAdpater := movable_adapater.New(obj)
	move := Move{}
	err := move.Execute(&movableAdpater)

	require.Equal(t, uobject.ErrNoDataByKey, err)
}

func TestMoveErrorGetVelocity(t *testing.T) {
	obj := uobject.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)

	movableAdpater := movable_adapater.New(obj)
	move := Move{}
	err = move.Execute(&movableAdpater)

	require.Equal(t, uobject.ErrNoDataByKey, err)
}

func TestMoveFreeze(t *testing.T) {
	obj := uobject.New()

	err := obj.SetProperty("Position", vector.Vector{12, 5})
	require.Equal(t, nil, err)
	err = obj.SetProperty("Direction", int(8))
	require.Equal(t, nil, err)
	err = obj.SetProperty("DirectionNumber", int(3))
	require.Equal(t, nil, err)
	err = obj.SetProperty("Velocity", int(10))
	require.Equal(t, nil, err)

	movableAdpater := movable_adapater_mock.New(obj)
	move := Move{}
	err = move.Execute(&movableAdpater)

	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{12, 5}, position)
}
