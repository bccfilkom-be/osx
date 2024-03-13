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

func BenchmarkWorkdir(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Workdir()
	}
}
