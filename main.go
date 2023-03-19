package main

import "math"

func SqueareRoot(a, b, c float64) []float64 {
	D := b*b - 4*a*c

	if D < 0 {
		return []float64{}
	}

	if D > 0 {
		return []float64{
			-b + math.Sqrt(D)/(2*a),
			-b - math.Sqrt(D)/(2*a),
		}
	}

	return nil
}
