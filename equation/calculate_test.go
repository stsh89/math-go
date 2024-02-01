package equation

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	got := Calculate("y=x+a", map[string]float64{"x": 1.0, "a": 5.0}, MockLogger{})
	want := 6.0

	if !floatsAreEqual(got, want) {
		t.Errorf("Calculate(\"y=x+a\") = %v, want %v", floatToString(got), floatToString(want))
	}
}

func TestCalculateRange(t *testing.T) {
	got := CalculateRange("y=x+a", map[string]float64{"a": 5.0}, MockLogger{})
	want := 11

	if len(got) != want {
		t.Errorf("CalculateRange(\"y=x+a\") = len(%v), want %v", len(got), want)
	}
}

func floatsAreEqual(a float64, b float64) bool {
	return floatToString(a) == floatToString(b)
}

func floatToString(value float64) string {
	return fmt.Sprintf("%.4f", value)
}
