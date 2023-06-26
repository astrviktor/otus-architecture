package interpret

import (
	"errors"
	"otus-architecture/hw11/broker"
	"otus-architecture/hw11/ioc"
	"otus-architecture/hw11/object"
)

type InterpretAdapter struct {
	*object.Object
}

func NewInterpretdapter(obj *object.Object) *InterpretAdapter {
	return &InterpretAdapter{obj}
}

func (a *InterpretAdapter) GetIoC() (*ioc.IoC, error) {
	return GetIoC(a.GetProperty("IoC"))
}

func (a *InterpretAdapter) GetMessage() (*broker.Message, error) {
	return GetMessage(a.GetProperty("Message"))
}

func GetIoC(i interface{}, err error) (*ioc.IoC, error) {
	if err != nil {
		return &ioc.IoC{}, err
	}

	value, ok := i.(*ioc.IoC)
	if !ok {
		return &ioc.IoC{}, ErrTypeCast
	}

	return value, nil
}

func GetMessage(i interface{}, err error) (*broker.Message, error) {
	if err != nil {
		return &broker.Message{}, err
	}

	value, ok := i.(*broker.Message)
	if !ok {
		return &broker.Message{}, ErrTypeCast
	}

	return value, nil
}

var (
	ErrTypeCast = errors.New("type cast error")
)
