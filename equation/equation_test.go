package equation

import (
	"testing"
)

type MockLogger struct{}

func (m MockLogger) Debug(description string) {}
func (m MockLogger) Error(description string) {}
func (m MockLogger) Info(description string)  {}
func (m MockLogger) Warn(description string)  {}

func TestDescription(t *testing.T) {
	got := Description()
	want := "Package provides operations around math equations"

	if got != want {
		t.Errorf("Description() = %v, want %v", got, want)
	}
}
