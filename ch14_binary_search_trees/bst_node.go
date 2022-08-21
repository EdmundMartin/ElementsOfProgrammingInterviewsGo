package ch14_binary_search_trees

type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(value int) *BSTNode {
	return &BSTNode{
		Data:  value,
		Left:  nil,
		Right: nil,
	}
}

func SearchBST(tree *BSTNode, key int) *BSTNode {
	if tree == nil {
		return tree
	} else if tree.Data == key {
		return tree
	} else if key < tree.Data {
		return SearchBST(tree.Left, key)
	} else {
		return SearchBST(tree.Right, key)
	}
}
