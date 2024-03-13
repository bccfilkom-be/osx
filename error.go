package osx

import "errors"

var (
	ErrUnrecoverableWorkdir = errors.New("osx: unable to get the current working directory")
)
