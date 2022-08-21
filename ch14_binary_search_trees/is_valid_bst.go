package ch14_binary_search_trees

import (
	"math"
)

func IsValidBST(tree *BSTNode) bool {
	lowRange := math.MinInt64
	highRange := math.MaxInt64
	return keysInRange(tree, lowRange, highRange)
}

func keysInRange(tree *BSTNode, lowRange, highRange int) bool {
	if tree == nil {
		return true
	} else if (tree.Data > lowRange && tree.Data <= highRange) == false {
		return false
	} else {
		return keysInRange(tree.Left, lowRange, tree.Data) && keysInRange(tree.Right, tree.Data, highRange)
	}
}

type queueEntry struct {
	Node  *BSTNode
	Lower int
	Upper int
}

func IsValidBSTIterative(tree *BSTNode) bool {

	bfsQueue := []queueEntry{
		{tree, math.MinInt64, math.MaxInt64},
	}

	for len(bfsQueue) > 0 {
		var front queueEntry
		front, bfsQueue = bfsQueue[0], bfsQueue[1:]

		if front.Node != nil {

			if front.Node.Data < front.Lower || front.Node.Data > front.Upper {
				return false
			}
			bfsQueue = append(bfsQueue, queueEntry{
				Node:  front.Node.Left,
				Lower: front.Lower,
				Upper: front.Node.Data,
			})

			bfsQueue = append(bfsQueue, queueEntry{
				Node:  front.Node.Right,
				Lower: front.Node.Data,
				Upper: front.Upper,
			})
		}
	}
	return true
}