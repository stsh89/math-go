package equation

import (
	"testing"
)

func TestDescription(t *testing.T) {
	got := Description()
	want := "Package provides operations around math equations"

	if got != want {
		t.Errorf("Description() = %v, want %v", got, want)
	}
}
