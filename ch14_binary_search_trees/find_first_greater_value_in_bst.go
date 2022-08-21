package ch14_binary_search_trees

/*
Write a program which takes a BST and a value and return the first key which would
appear in an inorder traversal which is greater than the input value.
*/

func FindFirstValueGreaterThanK(tree *BSTNode, k int) *BSTNode {
	subtree := tree
	var firstSoFar *BSTNode

	for subtree != nil {
		// If the subtree's data is greater than K - we want to search left
		if subtree.Data > k {
			firstSoFar, subtree = subtree, subtree.Left
		} else {
			// Can discard everything to lef as will always be smaller than K
			subtree = subtree.Right
		}
	}
	return firstSoFar
}
