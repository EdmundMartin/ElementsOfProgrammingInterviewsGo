package ch12_hash_tables

import (
	"testing"
)

func TestNearestRepeatedEntries(t *testing.T) {
	testArray := []string{
		"house",
		"cat",
		"dog",
		"mouse",
		"eagle",
		"house",
		"mouse",
	}

	result := NearestRepeatedEntries(testArray)

	if result != 3 {
		t.Error("wrong answer")
	}
}
