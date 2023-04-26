package getters

import (
	"errors"
	"otus-architecture/hw04/additions/vector"
)

func GetVector(i interface{}, err error) (vector.Vector, error) {
	if err != nil {
		return vector.Vector{}, err
	}

	value, ok := i.(vector.Vector)
	if !ok {
		return vector.Vector{}, ErrTypeCast
	}

	return value, nil
}

func GetInt(i interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	value, ok := i.(int)
	if !ok {
		return 0, ErrTypeCast
	}

	return value, nil
}

var (
	ErrTypeCast = errors.New("type cast error")
)
