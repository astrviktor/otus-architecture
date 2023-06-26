package getters

import (
	"errors"
	"otus-architecture/hw11/additions/queue"
	"otus-architecture/hw11/additions/vector"
)

func GetVector(i interface{}, err error) (vector.Vector, error) {
	if err != nil {
		return vector.Vector{}, err
	}

	value, ok := i.(vector.Vector)
	if !ok {
		return vector.Vector{}, ErrTypeCast
	}

	return value, nil
}

func GetAreaFunc(i interface{}, err error) (vector.Vector, error) {
	if err != nil {
		return vector.Vector{}, err
	}

	value, ok := i.(vector.Vector)
	if !ok {
		return vector.Vector{}, ErrTypeCast
	}

	return value, nil
}

func GetStringSlice(i interface{}, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}

	value, ok := i.([]string)
	if !ok {
		return nil, ErrTypeCast
	}

	return value, nil
}

func GetMapStringSlice(i interface{}, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}

	value, ok := i.([]string)
	if !ok {
		return nil, ErrTypeCast
	}

	return value, nil
}

func GetInt(i interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	value, ok := i.(int)
	if !ok {
		return 0, ErrTypeCast
	}

	return value, nil
}

//func GetString(i interface{}, err error) (string, error) {
//	if err != nil {
//		return "", err
//	}
//
//	value, ok := i.(string)
//	if !ok {
//		return "", ErrTypeCast
//	}
//
//	return value, nil
//}

//func GetIoC(i interface{}, err error) (ioc.IoC, error) {
//	if err != nil {
//		return ioc.IoC{}, err
//	}
//
//	value, ok := i.(ioc.IoC)
//	if !ok {
//		return ioc.IoC{}, ErrTypeCast
//	}
//
//	return value, nil
//}
//
//func GetMessage(i interface{}, err error) (broker.Message, error) {
//	if err != nil {
//		return broker.Message{}, err
//	}
//
//	value, ok := i.(broker.Message)
//	if !ok {
//		return broker.Message{}, ErrTypeCast
//	}
//
//	return value, nil
//}

func GetError(i interface{}, err error) (error, error) {
	if err != nil {
		return nil, err
	}

	value, ok := i.(error)
	if !ok {
		return nil, ErrTypeCast
	}

	return value, nil
}

func GetStructChannel(i interface{}, err error) (chan struct{}, error) {
	if err != nil {
		return nil, err
	}

	ch, ok := i.(chan struct{})
	if !ok {
		return nil, ErrTypeCast
	}

	return ch, nil
}

func GetQueueChannel(i interface{}, err error) (chan queue.Element, error) {
	if err != nil {
		return nil, err
	}

	ch, ok := i.(chan queue.Element)
	if !ok {
		return nil, ErrTypeCast
	}

	return ch, nil
}

var (
	ErrTypeCast = errors.New("type cast error")
)
