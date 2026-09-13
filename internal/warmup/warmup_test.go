package warmup

import (
	"errors"
	"testing"
)

var TestName = "Director"

func TestByteOffset(t *testing.T) {
	w := NewWriter(TestName)

	if w.Name != TestName {
		t.Fatal("Expected name as 'Director', instead got", w.Name)
	}

	w.MoveOffset(10)

	if w.offset != 10 {
		t.Fatal("Expected 10, Got", w.offset)
	}

	w.MoveOffset(27)

	if w.offset != 37 {
		t.Fatal("Expected 37, Got", w.offset)
	}

	err := w.MoveOffset(-14)

	if !errors.Is(err, ErrNegativeDelta) {
		t.Errorf("Expected %s, got %s", ErrNegativeDelta, err)
	}
}
