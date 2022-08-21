package ch11_searching

func SearchShiftedSortedSmallest(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		middle := (left + right) / 2
		if arr[middle] > arr[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
