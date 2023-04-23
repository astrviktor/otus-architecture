package uobject

import (
	"errors"
)

type UObjectInterface interface {
	GetProperty(key string) (interface{}, error)
	SetProperty(key string, object interface{}) error
}

type UObject struct {
	Property map[string]interface{}
}

func New() *UObject {
	return &UObject{
		Property: make(map[string]interface{}),
	}
}

func (o *UObject) GetProperty(key string) (interface{}, error) {
	object, ok := o.Property[key]
	if !ok {
		return nil, ErrNoDataByKey
	}

	return object, nil
}

func (o *UObject) SetProperty(key string, object interface{}) error {
	o.Property[key] = object
	return nil
}

var (
	ErrNoDataByKey = errors.New("no data by key")
)
