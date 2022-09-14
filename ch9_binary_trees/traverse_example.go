package ch9_binary_trees

import "fmt"

func TreeTraverse(root *BinaryTreeNode) {
	if root != nil {
		fmt.Sprintf("Pre-Order: %d", root.Data)
		TreeTraverse(root.Left)
		fmt.Sprintf("In-Order: %d", root.Data)
		TreeTraverse(root.Right)
		fmt.Sprintf("Postorder: %d", root.Data)
	}
}
