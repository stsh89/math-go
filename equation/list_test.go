package equation

import (
	"testing"
)

type MockLister struct{}

func (m MockLister) List() []string {
	return []string{
		"y=sinx",
		"y=ab+x",
	}
}

func TestList(t *testing.T) {
	lister := MockLister{}

	got := List(lister, MockLogger{})

	if len(got) != 2 {
		t.Errorf("List(lister, logger) = length %v, want length %v", len(got), 2)
	}
}
