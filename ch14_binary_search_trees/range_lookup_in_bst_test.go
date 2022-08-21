package ch14_binary_search_trees

import (
	"reflect"
	"testing"
)

func TestRangeLookupInBST(t *testing.T) {
	root := BuildMinHeightSortedArray([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120})
	results := RangeLookupInBST(root, Interval{
		Left:  20,
		Right: 70,
	})
	if !reflect.DeepEqual(results, []int{20, 30, 40, 50, 60, 70}) {
		t.Error("unexpected result")
	}
}
