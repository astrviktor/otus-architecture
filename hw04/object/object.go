package object

import "errors"

type InterfaceObject interface {
	GetProperty(key string) (interface{}, error)
	SetProperty(key string, object interface{}) error
}

type Object struct {
	Property map[string]interface{}
}

func New() *Object {
	return &Object{
		Property: make(map[string]interface{}),
	}
}

func (o *Object) GetProperty(key string) (interface{}, error) {
	object, ok := o.Property[key]
	if !ok {
		return nil, ErrNoDataByKey
	}

	return object, nil
}

func (o *Object) SetProperty(key string, object interface{}) error {
	o.Property[key] = object
	return nil
}

var (
	ErrNoDataByKey = errors.New("no data by key")
)
