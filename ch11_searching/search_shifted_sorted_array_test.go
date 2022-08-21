package ch11_searching

import "testing"

func TestSearchShiftedSortedSmallest(t *testing.T) {
	arr := []int{378, 478, 550, 631, 103, 203, 220, 234, 279, 368}
	result := SearchShiftedSortedSmallest(arr)
	if arr[result] != 103 {
		t.Error("expected to get smallest value")
	}
}
