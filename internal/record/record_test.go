package record

import (
	"testing"
)

func TestHeaderSize(t *testing.T) {
	if HeaderSize != 18 {
		t.Error("Expected Header size to be 18, instead got:", HeaderSize)
	}
}
