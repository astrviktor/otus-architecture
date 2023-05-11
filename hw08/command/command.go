package command

import "otus-architecture/hw08/object"

type CommandInterface interface {
	Execute() (*object.Object, error)
}
