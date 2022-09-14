package ch8_stacks_queues

import "testing"

func TestIsWellFormed(t *testing.T) {
	badlyFormed := "[(])"

	if IsWellFormed(badlyFormed) {
		t.Error("should be badly formed")
	}

	wellFormed := "[({})]"
	if !IsWellFormed(wellFormed) {
		t.Error("should be well formed")
	}
}
