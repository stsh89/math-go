package equation

import (
	"testing"
)

type MockFinder struct{}

func (m MockFinder) Find(term string) string {
	return ""
}

func TestFind(t *testing.T) {
	finder := MockFinder{}

	got := Find("y=sinx", finder, MockLogger{})
	want := ""

	if got != want {
		t.Errorf("Save(\"y=sinx\", saver) = %v, want %v", got, want)
	}
}
