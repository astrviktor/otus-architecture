package iocstrategies

import "otus-architecture/hw08/object"

type Nop struct{}

func (c *Nop) Execute() (*object.Object, error) {
	return nil, nil
}
