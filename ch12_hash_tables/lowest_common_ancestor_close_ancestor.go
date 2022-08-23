package ch12_hash_tables

type BinaryTreeNodeWithParent struct {
	Data   int
	Left   *BinaryTreeNodeWithParent
	Right  *BinaryTreeNodeWithParent
	Parent *BinaryTreeNodeWithParent
}

func LCA(firstNode, secondNode *BinaryTreeNodeWithParent) *BinaryTreeNodeWithParent {
	firstIter, secondIter := firstNode, secondNode
	nodesOnPath := make(map[*BinaryTreeNodeWithParent]interface{})

	for firstIter != nil || secondIter != nil {

		if firstIter != nil {
			_, ok := nodesOnPath[firstIter]
			if ok {
				return firstIter
			}
			nodesOnPath[firstIter] = nil
			firstIter = firstIter.Parent
		}

		if secondIter != nil {
			_, ok := nodesOnPath[secondIter]
			if ok {
				return secondIter
			}
			nodesOnPath[secondIter] = nil
			secondIter = secondIter.Parent
		}
	}
	// Given nodes are in different trees
	return nil
}