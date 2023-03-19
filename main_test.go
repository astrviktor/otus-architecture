package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSqueareRoot(t *testing.T) {
	tests := []struct {
		a, b, c  float64
		expected []float64
	}{
		{1, 0, 1, []float64{}},
		{1, 0, -1, []float64{1, -1}},
		{1, 2, 1, []float64{-1}},
	}

	for _, test := range tests {
		result := SqueareRoot(test.a, test.b, test.c)
		require.Equal(t, test.expected, result)
	}
}
