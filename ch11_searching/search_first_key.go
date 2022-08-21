package ch11_searching

func SearchFirstKey(arr []int, k int) int {
	result := -1
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2
		if arr[middle] > k {
			right = middle - 1
		} else if arr[middle] == k {
			result = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return result
}