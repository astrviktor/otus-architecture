package matrix

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestWriteMatrixToFile(t *testing.T) {
	const FileName = "F1.txt"
	const CountColumns = 3
	const CountRows = 2

	matrix := GenerateMatrix(CountColumns, CountRows)

	err := WriteMatrixToFile(FileName, matrix)
	require.Equal(t, nil, err)

	err = os.Remove(FileName)
	require.Equal(t, nil, err)
}

func TestReadMatrixFromFile(t *testing.T) {
	const FileName = "F2.txt"
	const CountColumns = 3
	const CountRows = 2

	matrix1 := GenerateMatrix(CountColumns, CountRows)

	err := WriteMatrixToFile(FileName, matrix1)
	require.Equal(t, nil, err)

	matrix2 := GenerateMatrix(CountColumns, CountRows)

	err = WriteMatrixToFile(FileName, matrix2)
	require.Equal(t, nil, err)

	matrixes, err := ReadMatrixFromFile(FileName)
	require.Equal(t, nil, err)

	require.Equal(t, 2, len(matrixes))

	require.Equal(t, matrix1, matrixes[0])
	require.Equal(t, matrix2, matrixes[1])

	err = os.Remove(FileName)
	require.Equal(t, nil, err)
}
