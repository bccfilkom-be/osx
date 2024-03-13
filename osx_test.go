package osx

import (
	"testing"
)

func TestWorkdir(t *testing.T) {
	wd, err := Workdir()
	if err != nil {
		t.Fatalf("expected: %v\nactual: %s", nil, err)
	}

	if len(wd) == 0 {
		t.Errorf("expected: %v\nactual: %s", "not empty", wd)
	}
}
