package command

import "otus-architecture/hw13/object"

type CommandInterface interface {
	Execute() (*object.Object, error)
}
