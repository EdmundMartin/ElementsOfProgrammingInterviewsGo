package ch11_searching

import "testing"

func TestSearchEntryEqualToIndex(t *testing.T) {

	hasMatchingIndex := []int{-10, -8, -7, 1, 4, 50, 60}

	result := SearchEntryEqualToIndex(hasMatchingIndex)

	if hasMatchingIndex[result] != 4 {
		t.Error("unexpected value")
	}

	noMatchingIndex := []int{10, 20, 30, 40, 50}

	if SearchEntryEqualToIndex(noMatchingIndex) != -1 {
		t.Error("expected value to -1")
	}
}
