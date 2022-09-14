package ch9_binary_trees

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTreeNodeWithParent struct {
	Data   int
	Left   *BinaryTreeNodeWithParent
	Right  *BinaryTreeNodeWithParent
	Parent *BinaryTreeNodeWithParent
}
