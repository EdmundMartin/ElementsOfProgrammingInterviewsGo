package ch14_binary_search_trees

func RebuildBSTFromPreOrder(preOrderSeq []int) *BSTNode {
	if len(preOrderSeq) == 0 {
		return nil
	}

	if len(preOrderSeq) == 1 {
		return &BSTNode{
			Data:  preOrderSeq[0],
			Left:  nil,
			Right: nil,
		}
	}

	var transitionPoint int
	for idx, val := range preOrderSeq {
		if val > preOrderSeq[0] {
			transitionPoint = idx
			break
		}
	}

	return &BSTNode{
		Data:  preOrderSeq[0],
		Left:  RebuildBSTFromPreOrder(preOrderSeq[1:transitionPoint]),
		Right: RebuildBSTFromPreOrder(preOrderSeq[transitionPoint:]),
	}
}
