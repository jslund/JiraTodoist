package function

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	input := "TestTask"

	err := AddItem(input, 1)
	if err != nil {
		t.Fatalf("Couldn't add task")
	}
}
