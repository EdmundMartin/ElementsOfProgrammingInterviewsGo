package ch11_searching

func MatrixSearch(arr [][]int, target int) bool {
	row, col := 0, len(arr[0])

	// We start in the top right corner of the matrix
	for row < len(arr) && col >= 0 {
		if arr[row][col] == target {
			return true
		} else if arr[row][col] < target {
			row += 1
		} else {
			col -= 1
		}
	}
	return false
}
