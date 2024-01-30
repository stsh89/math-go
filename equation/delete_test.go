package equation

import (
	"testing"
)

type MockDeleter struct {
	id string
}

func (m MockDeleter) Delete(term string) string {
	return m.id
}

func TestDelete(t *testing.T) {
	deleter := buildDeleter()

	got := Delete("y=sinx", deleter, MockLogger{})
	want := "cfec7989-fa56-415e-9b82-9fa7693f3dcf"

	if got != want {
		t.Errorf("Save(\"y=sinx\", deleter) = %v, want %v", got, want)
	}
}

func buildDeleter() MockDeleter {
	return MockDeleter{
		id: "cfec7989-fa56-415e-9b82-9fa7693f3dcf",
	}
}
