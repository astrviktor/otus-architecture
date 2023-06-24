package iocstrategies

import "otus-architecture/hw11/object"

type Nop struct{}

func (c *Nop) Execute() (*object.Object, error) {
	return nil, nil
}
