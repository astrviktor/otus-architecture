package command

import "otus-architecture/hw11/object"

type CommandInterface interface {
	Execute() (*object.Object, error)
}
