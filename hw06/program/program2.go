package program

import (
	"otus-architecture/hw06/adapter"
	"otus-architecture/hw06/matrix"
)

type InterfaceProgram2 interface {
	Write2MatrixToFile(name string) error
	ExecuteMatrixAdd(input, output string) error
}

type Program2 struct {
}

func (p *Program2) Write2MatrixToFile(name string) error {
	const CountColumns = 3
	const CountRows = 2

	matrix1 := matrix.GenerateMatrix(CountColumns, CountRows)
	matrix2 := matrix.GenerateMatrix(CountColumns, CountRows)

	err := matrix.WriteMatrixToFile(name, matrix1)
	if err != nil {
		return err
	}

	err = matrix.WriteMatrixToFile(name, matrix2)
	if err != nil {
		return err
	}

	return nil
}

func (p *Program2) ExecuteMatrixAdd(input, output string) error {
	e := adapter.Executor{}
	a := adapter.NewSummatorAdapater(&e)
	err := a.ExecuteSummator(input, output)

	return err
}
