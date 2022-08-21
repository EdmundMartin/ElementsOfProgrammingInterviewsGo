package ch11_searching

import "testing"

func TestFindMinMax(t *testing.T) {

	testArr := []int{3, 2, 5, 1, 2, 4}

	result := FindMinMax(testArr)

	if result.Smallest != 1 {
		t.Error("not smallest value")
	}

	if result.Largest != 5 {
		t.Error("not largest value")
	}
}
