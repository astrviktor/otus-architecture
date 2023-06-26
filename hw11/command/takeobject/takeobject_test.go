package takeobject

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw11/command/makeobject"
	"otus-architecture/hw11/object"
	"testing"
)

func TestTakeObjectOk(t *testing.T) {
	const Game = 123
	const Object = 456

	obj := object.New()

	obj.SetProperty("GameIdentifier", Game)
	obj.SetProperty("ObjectIdentifier", Object)

	makeObjectAdapter := makeobject.NewMakeObjectAdapter(obj)
	makeObjectCommand := makeobject.NewMakeObjectCommand(makeObjectAdapter)
	_, err := makeObjectCommand.Execute()

	takeObjectAdapter := NewTakeObjectAdapter(obj)
	takeObjectCommand := NewTakeObjectCommand(takeObjectAdapter)
	takeObject, err := takeObjectCommand.Execute()

	require.Equal(t, nil, err)
	require.Equal(t, obj, takeObject)

	gameIdentifier, err := takeObject.GetProperty("GameIdentifier")
	require.Equal(t, Game, gameIdentifier)
	objectIdentifier, err := takeObject.GetProperty("ObjectIdentifier")
	require.Equal(t, Object, objectIdentifier)
}
