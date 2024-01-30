package equation

import (
	"testing"
)

type MockSaver struct {
	id string
}

func (m MockSaver) Save(term string) string {
	return m.id
}

func TestSave(t *testing.T) {
	saver := buildSaver()

	got := Save("y=sinx", saver, MockLogger{})
	want := "407b46bf-9630-4257-b075-a0ca8f9ed7b7"

	if got != want {
		t.Errorf("Save(\"y=sinx\", saver) = %v, want %v", got, want)
	}
}

func buildSaver() MockSaver {
	return MockSaver{
		id: "407b46bf-9630-4257-b075-a0ca8f9ed7b7",
	}
}
