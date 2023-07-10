package iocstrategies

import "otus-architecture/hw13/object"

type Nop struct{}

func (c *Nop) Execute() (*object.Object, error) {
	return nil, nil
}
