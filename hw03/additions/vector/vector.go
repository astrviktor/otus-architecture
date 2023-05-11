package vector

type Vector struct {
	X, Y int
}

func VectorPlus(v1, v2 Vector) Vector {
	var v Vector

	v.X = v1.X + v2.X
	v.Y = v1.Y + v2.Y

	return v
}
