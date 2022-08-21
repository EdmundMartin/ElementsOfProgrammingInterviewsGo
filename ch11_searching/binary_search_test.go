package ch11_searching

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	result := BinarySearch(arr, 30)
	fmt.Println(result)

	if result != 2 {
		t.Error("30 is at index 2")
	}

	result = BinarySearch(arr, 120)

	if result != -1 {
		t.Error("120 is not in array")
	}
}