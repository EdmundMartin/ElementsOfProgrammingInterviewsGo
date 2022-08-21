package ch14_binary_search_trees

func FindKLargestInBST(tree *BSTNode, k int) []int {
	var values []int
	findKHelper(tree, k, &values)
	return values
}

func findKHelper(tree *BSTNode, k int, values *[]int) {
	if tree != nil && len(*values) < k {
		findKHelper(tree.Right, k, values)
		if len(*values) < k {
			*values = append(*values, tree.Data)
			findKHelper(tree.Left, k, values)
		}
	}
}
