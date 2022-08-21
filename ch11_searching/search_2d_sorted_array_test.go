package ch11_searching

import "testing"

func TestMatrixSearch(t *testing.T) {
	results := [][]int{
		{-1, 2, 4, 4, 6},
		{1, 5, 5, 9, 21},
		{3, 6, 6, 9, 22},
		{3, 6, 8, 10, 24},
		{6, 8, 9, 12, 25},
	}

	found := MatrixSearch(results, 9)
	if found != true {
		t.Error("9 is in the matrix")
	}
}
