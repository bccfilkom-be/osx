package osx

import (
	"path"
	"runtime"
)

func Workdir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", ErrUnrecoverableWorkdir
	}
	dir := path.Dir(filename)
	return dir, nil
}
