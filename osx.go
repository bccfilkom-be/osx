package osx

import (
	"path"
	"runtime"
)

// Get current working directory by recovering the information of the file caller.
// After getting the filename, retrieve the corresponding directory it live and
// then return the absolute path of that directory. It will return an error if it
// fails to retrieve the filename.
func Workdir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", ErrUnrecoverableWorkdir
	}
	dir := path.Dir(filename)
	return dir, nil
}
