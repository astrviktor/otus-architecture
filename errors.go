package main

import "errors"

var (
	ErrAIsZero    = errors.New("A is zero")
	ErrUnexpected = errors.New("unexpected result")
)
