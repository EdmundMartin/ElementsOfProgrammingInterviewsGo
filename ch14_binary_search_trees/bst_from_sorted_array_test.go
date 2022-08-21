package ch14_binary_search_trees

import (
	"testing"
)

func TestBuildMinHeightSortedArray(t *testing.T) {

	sortedArr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	result := BuildMinHeightSortedArray(sortedArr)

	if result.Data != 60 {
		t.Error("expected 60 as middle index")
	}

	if result.Left.Data != 30 {
		t.Error("expected left child to be 30")
	}
}
