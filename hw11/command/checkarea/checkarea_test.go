package checkarea

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw11/additions/vector"
	"otus-architecture/hw11/object"
	"testing"
)

// функции получения координат окресности из Position
func base(position vector.Vector) vector.Vector {
	return vector.Vector{X: position.X / 10, Y: position.Y / 10}
}

func basePlus5(position vector.Vector) vector.Vector {
	return vector.Vector{X: (position.X + 5) / 10, Y: (position.Y + 5) / 10}
}

func basePlus7(position vector.Vector) vector.Vector {
	return vector.Vector{X: (position.X + 7) / 10, Y: (position.Y + 7) / 10}
}

func TestCheckAreaOk(t *testing.T) {
	AreasNames := []string{"base", "base+5", "base+7"}
	AreasFuncs := map[string]func(vector.Vector) vector.Vector{
		"base":   base,
		"base+5": basePlus5,
		"base+7": basePlus7,
	}

	obj1 := object.New()

	obj1.SetProperty("AreasNames", AreasNames)
	obj1.SetProperty("AreasFuncs", AreasFuncs)
	obj1.SetProperty("Position", vector.Vector{10, 10})

	obj2 := object.New()

	obj2.SetProperty("AreasNames", AreasNames)
	obj2.SetProperty("AreasFuncs", AreasFuncs)
	obj2.SetProperty("Position", vector.Vector{15, 15})

	checkAreaAdapter1 := NewCheckAreaAdapter(obj1)
	checkAreaCommand1 := NewCheckAreaCommand(checkAreaAdapter1)

	checkAreaAdapter2 := NewCheckAreaAdapter(obj2)
	checkAreaCommand2 := NewCheckAreaCommand(checkAreaAdapter2)

	_, err := checkAreaCommand1.Execute()
	require.Equal(t, nil, err)

	_, err = checkAreaCommand2.Execute()
	require.Equal(t, nil, err)
}
