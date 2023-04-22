package hw01

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSqueareRoot(t *testing.T) {
	tests := []struct {
		a, b, c  float64
		expected []float64
		err      error
	}{
		{1, 0, 1, []float64{}, nil},
		{1, 0, -1, []float64{1, -1}, nil},
		{1, 2, 1, []float64{-1}, nil},
		{0, 1, 2, nil, ErrAIsZero},
	}

	for _, test := range tests {
		result, err := SqueareRoot(test.a, test.b, test.c)
		require.Equal(t, test.expected, result)
		require.Equal(t, test.err, err)
	}
}
