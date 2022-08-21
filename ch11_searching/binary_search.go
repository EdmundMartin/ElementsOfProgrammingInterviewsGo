package ch11_searching

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] < target {
			left = middle + 1
		} else if arr[middle] == target {
			return middle
		} else {
			right = middle - 1
		}
	}
	return -1
}
