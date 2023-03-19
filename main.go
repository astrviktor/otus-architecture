package main

func SqueareRoot(a, b, c float64) []float64 {
	D := b*b - 4*a*c
	if D < 0 {
		return []float64{}
	}
	return nil
}
