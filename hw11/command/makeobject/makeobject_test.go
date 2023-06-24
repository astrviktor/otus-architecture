package makeobject

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw11/object"
	"testing"
)

func TestMakeObjectOk(t *testing.T) {
	const Game = 123
	const Object = 456

	obj := object.New()

	obj.SetProperty("GameIdentifier", Game)
	obj.SetProperty("ObjectIdentifier", Object)

	makeObjectAdapter := NewMakeObjectAdapter(obj)
	makeObjectCommand := NewMakeObjectCommand(makeObjectAdapter)
	resObj, err := makeObjectCommand.Execute()

	require.Equal(t, nil, err)
	require.Equal(t, resObj, obj)

	gameIdentifier, err := obj.GetProperty("GameIdentifier")
	require.Equal(t, Game, gameIdentifier)
	objectIdentifier, err := obj.GetProperty("ObjectIdentifier")
	require.Equal(t, Object, objectIdentifier)

	require.Equal(t, 1, len(object.Games))
	require.Equal(t, 1, len(object.Games[Game]))

	storeObj, ok := object.Games[Game][Object]
	require.Equal(t, true, ok)
	require.Equal(t, storeObj, obj)
}
