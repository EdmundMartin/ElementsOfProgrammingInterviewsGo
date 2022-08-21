package ch14_binary_search_trees

func BuildMinHeightSortedArray(arr []int) *BSTNode {
	return buildMinHeightBSTFromSortedSubArray(arr, 0, len(arr))
}

func buildMinHeightBSTFromSortedSubArray(arr []int, start, end int) *BSTNode {
	if start >= end {
		return nil
	}
	middle := (start + end) / 2

	return &BSTNode{
		Data:  arr[middle],
		Left:  buildMinHeightBSTFromSortedSubArray(arr, start, middle),
		Right: buildMinHeightBSTFromSortedSubArray(arr, middle+1, end),
	}
}
