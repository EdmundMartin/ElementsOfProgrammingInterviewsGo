package ch11_searching

import "testing"

func TestSearchFirstKey(t *testing.T) {
	arr := []int{10, 20, 30, 30, 30, 40, 50, 60, 70, 110}

	idx := SearchFirstKey(arr, 30)

	if idx != 2 {
		t.Error("got wrong idx")
	}
}
