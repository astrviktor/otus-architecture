package matrix

import (
	"encoding/json"
	"errors"
	"math/rand"
)

type Matrix struct {
	CountColumns int
	CountRows    int
	Data         [][]int
}

func GenerateMatrix(CountColumns, CountRows int) *Matrix {
	const MIN = -100
	const MAX = 100

	data := make([][]int, CountRows)

	for r := 0; r < CountRows; r++ {

		row := make([]int, CountColumns)
		for c := 0; c < CountColumns; c++ {
			row[c] = rand.Intn(MAX-MIN+1) + MIN
		}

		data[r] = row
	}

	return &Matrix{
		CountColumns: CountColumns,
		CountRows:    CountRows,
		Data:         data,
	}
}

func MatrixToBytes(m *Matrix) ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func BytesToMatrix(bytes []byte) (*Matrix, error) {
	matrix := Matrix{}
	err := json.Unmarshal(bytes, &matrix)
	if err != nil {
		return nil, err
	}

	return &matrix, nil
}

func MatrixAdd(m1, m2 *Matrix) (*Matrix, error) {
	if m1.CountColumns != m2.CountColumns {
		return nil, ErrMatrixSizeNotEqual
	}

	if m1.CountRows != m2.CountRows {
		return nil, ErrMatrixSizeNotEqual
	}

	CountColumns := m1.CountColumns
	CountRows := m1.CountRows

	data := make([][]int, CountRows)

	for r := 0; r < CountRows; r++ {

		row := make([]int, CountColumns)
		for c := 0; c < CountColumns; c++ {
			row[c] = m1.Data[r][c] + m2.Data[r][c]
		}

		data[r] = row
	}

	result := Matrix{
		CountColumns: CountColumns,
		CountRows:    CountRows,
		Data:         data,
	}

	return &result, nil
}

var (
	ErrMatrixSizeNotEqual = errors.New("matrix size not equal")
)
