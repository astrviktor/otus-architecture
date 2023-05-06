package matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	const CountColumns = 5
	const CountRows = 6

	matrix := GenerateMatrix(CountColumns, CountRows)

	require.Equal(t, CountRows, len(matrix.Data))
	require.Equal(t, CountColumns, len(matrix.Data[0]))
}

func TestMatrixMarshalAndUnmarshal(t *testing.T) {
	const CountColumns = 5
	const CountRows = 6

	matrix1 := GenerateMatrix(CountColumns, CountRows)

	bytes, err := MatrixToBytes(matrix1)
	require.Equal(t, nil, err)

	//fmt.Println(string(bytes))

	matrix2, err := BytesToMatrix(bytes)
	require.Equal(t, nil, err)

	require.Equal(t, matrix1, matrix2)
}

func TestAddMatrixOk(t *testing.T) {
	bytesMatrix1 := []byte(`{"CountColumns":3,"CountRows":2,"Data":[[1,2,3],[4,5,6]]}`)
	bytesMatrix2 := []byte(`{"CountColumns":3,"CountRows":2,"Data":[[6,5,4],[3,2,1]]}`)
	bytesMatrixExpected := []byte(`{"CountColumns":3,"CountRows":2,"Data":[[7,7,7],[7,7,7]]}`)

	matrix1, err := BytesToMatrix(bytesMatrix1)
	require.Equal(t, nil, err)

	matrix2, err := BytesToMatrix(bytesMatrix2)
	require.Equal(t, nil, err)

	matrixExpected, err := BytesToMatrix(bytesMatrixExpected)
	require.Equal(t, nil, err)

	matrixResult, err := MatrixAdd(matrix1, matrix2)
	require.Equal(t, nil, err)

	require.Equal(t, matrixExpected, matrixResult)
}

func TestAddMatrixError(t *testing.T) {
	bytesMatrix1 := []byte(`{"CountColumns":2,"CountRows":2,"Data":[[1,2],[3,4]]}`)
	bytesMatrix2 := []byte(`{"CountColumns":3,"CountRows":2,"Data":[[6,5,4],[3,2,1]]}`)

	matrix1, err := BytesToMatrix(bytesMatrix1)
	require.Equal(t, nil, err)

	matrix2, err := BytesToMatrix(bytesMatrix2)
	require.Equal(t, nil, err)

	_, err = MatrixAdd(matrix1, matrix2)
	require.Equal(t, ErrMatrixSizeNotEqual, err)
}
