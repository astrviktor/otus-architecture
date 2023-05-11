package matrix

import (
	"os"
	"strings"
)

func WriteMatrixToFile(name string, matrix *Matrix) error {
	bytes, err := MatrixToBytes(matrix)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(bytes) + "\n")

	return err
}

func ReadMatrixFromFile(name string) ([]*Matrix, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	var result []*Matrix

	for _, line := range lines {
		if line == "" {
			continue
		}

		matrix, err := BytesToMatrix([]byte(line))

		if err != nil {
			return nil, err
		}

		result = append(result, matrix)
	}

	return result, nil
}
