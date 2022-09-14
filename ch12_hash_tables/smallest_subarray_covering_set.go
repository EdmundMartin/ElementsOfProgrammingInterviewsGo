package ch12_hash_tables

type SubArrayInfo struct {
	Start int
	End   int
}

func (s SubArrayInfo) Equal(other SubArrayInfo) bool {
	return s.Start == other.Start && s.End == other.End
}

func buildKeyWordCounter(keywords map[string]interface{}) map[string]int {
	result := make(map[string]int)
	for key, _ := range keywords {
		_, ok := result[key]
		if !ok {
			result[key] = 1
		} else {
			result[key]++
		}
	}
	return result
}

func FindSmallestSubArrayCoveringSet(paragraph []string, keywords map[string]interface{}) SubArrayInfo {
	keywordsToCover := buildKeyWordCounter(keywords)
	result := SubArrayInfo{-1, -1}
	remainingToCover := len(keywords)
	left := 0

	for right, p := range paragraph {
		_, ok := keywords[p]
		if ok {
			keywordsToCover[p] -= 1
			if keywordsToCover[p] >= 0 {
				remainingToCover -= 1
			}
		}

		for remainingToCover == 0 {
			if result.Equal(SubArrayInfo{-1, -1}) || (right-left) < (result.End-result.Start) {
				result = SubArrayInfo{left, right}
			}

			pl := paragraph[left]
			if _, ok := keywords[pl]; ok {
				keywordsToCover[pl] += 1
				if keywordsToCover[p] > 0 {
					remainingToCover += 1
				}
			}
			left += 1
		}
	}
	return result
}

type DoublyLinkedListNode struct {
	Data int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

func NewDoublyLinkedNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Data: val,
		Next: nil,
		Prev: nil,
	}
}

type LinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Len() int {
	return l.Size
}

func (l *LinkedList) InsertAfter(value int) {
	node := NewDoublyLinkedNode(value)
	node.Prev = l.Tail
	if l.Tail != nil {
		l.Tail.Next = node
	} else {
		l.Head = node
	}
	l.Tail = node
	l.Size++
}

func (l *LinkedList) Remove(node *DoublyLinkedListNode) {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}
	node.Next = nil
	node.Prev = nil
	l.Size--
}

func FindSmallestSubArrayCoveringSetLinkedList(paragraph []string, keywords []string) SubArrayInfo {
	loc := NewLinkedList()

	mapping := make(map[string]*DoublyLinkedListNode)
	result := SubArrayInfo{-1, -1}

	for idx, s := range paragraph {
		if _, ok := mapping[s]; ok {
			val := mapping[s]
			if val != nil {
				loc.Remove(val)
			}
			loc.InsertAfter(idx)
			mapping[s] = loc.Tail

			if loc.Size == len(keywords) {
				if result.Equal(SubArrayInfo{-1, -1}) || (idx-loc.Head.Data) < (result.End-result.Start) {
					result = SubArrayInfo{loc.Head.Data, idx}
				}
			}
		}
	}
	return result
}
