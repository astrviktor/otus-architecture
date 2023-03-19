package main

import (
	"math"
)

func SqueareRoot(a, b, c float64) ([]float64, error) {
	const epsilon = 0.0000001
	D := b*b - 4*a*c

	if math.Abs(a) < epsilon {
		return nil, ErrAIsZero
	}

	if math.Abs(D) < epsilon {
		return []float64{
			-b / (2 * a),
		}, nil
	}

	if D < 0 {
		return []float64{}, nil
	}

	if D > 0 {
		return []float64{
			-b + math.Sqrt(D)/(2*a),
			-b - math.Sqrt(D)/(2*a),
		}, nil
	}

	return nil, ErrUnexpected
}
