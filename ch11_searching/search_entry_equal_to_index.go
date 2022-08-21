package ch11_searching

func SearchEntryEqualToIndex(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := (left + right) / 2

		difference := arr[middle] - middle
		if difference == 0 {
			return middle
		} else if difference > 0 {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
