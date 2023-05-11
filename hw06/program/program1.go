package program

import (
	"errors"
	"otus-architecture/hw06/matrix"
)

type InterfaceProgram1 interface {
	Read2MatrixFromFile(name string) ([]*matrix.Matrix, error)
	MatrixAdd(m1, m2 *matrix.Matrix) (*matrix.Matrix, error)
	WriteResultMatrixToFile(name string, result *matrix.Matrix) error
	ExecuteMatrixAdd(input, output string) error
}

type Program1 struct{}

func (p *Program1) Read2MatrixFromFile(name string) ([]*matrix.Matrix, error) {
	return matrix.ReadMatrixFromFile(name)
}

func (p *Program1) MatrixAdd(m1, m2 *matrix.Matrix) (*matrix.Matrix, error) {
	return matrix.MatrixAdd(m1, m2)
}

func (p *Program1) WriteResultMatrixToFile(name string, result *matrix.Matrix) error {
	return matrix.WriteMatrixToFile(name, result)
}

func (p *Program1) ExecuteMatrixAdd(input, output string) error {
	matrixes, err := p.Read2MatrixFromFile(input)
	if err != nil {
		return err
	}

	if len(matrixes) < 2 {
		return errors.New("required 2 matrixes")
	}

	result, err := p.MatrixAdd(matrixes[0], matrixes[1])
	if err != nil {
		return err
	}

	err = p.WriteResultMatrixToFile(output, result)

	return err
}
