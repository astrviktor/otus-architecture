package exception

import (
	"otus-architecture/hw12/command"
	"otus-architecture/hw12/object"
	"reflect"
)

type HandlerFunc func(cmd command.CommandInterface, err error)
type HandlerError map[error]HandlerFunc

type Exception struct {
	exceptions map[string]HandlerError
}

func NewException() *Exception {
	exceptions := map[string]HandlerError{}

	return &Exception{
		exceptions: exceptions,
	}
}

func NewExceptionRepeatLog() *Exception {
	exceptions := map[string]HandlerError{
		"*move.MoveCommand":     {object.ErrNoDataByKey: AddRepeatCommand},
		"*repeat.RepeatCommand": {object.ErrNoDataByKey: AddLogCommand},
	}

	return &Exception{
		exceptions: exceptions,
	}
}

func NewExceptionRepeatRepeatLog() *Exception {
	exceptions := map[string]HandlerError{
		"*move.MoveCommand":               {object.ErrNoDataByKey: AddRepeatCommand},
		"*repeat.RepeatCommand":           {object.ErrNoDataByKey: AddRepeatCountCommand},
		"*repeatcount.RepeatCountCommand": {object.ErrNoDataByKey: AddLogCommand},
	}

	return &Exception{
		exceptions: exceptions,
	}
}

func (e *Exception) Handle(cmd command.CommandInterface, err error) {
	commandType := reflect.TypeOf(cmd).String()

	handlerError, ok := e.exceptions[commandType]
	if !ok {
		return
	}

	handlerFunc, ok := handlerError[err]
	if !ok {
		return
	}

	handlerFunc(cmd, err)

	return
}
