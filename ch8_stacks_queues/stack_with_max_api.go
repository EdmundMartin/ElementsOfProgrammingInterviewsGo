package ch8_stacks_queues

type ElementWithCachedMax struct {
	Element int
	Max     int
}

type StackWithMax struct {
	elements []ElementWithCachedMax
}

func NewStackWithMax() *StackWithMax {
	return &StackWithMax{elements: []ElementWithCachedMax{}}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *StackWithMax) Empty() bool {
	return len(s.elements) == 0
}

func (s *StackWithMax) Max() int {
	return s.elements[len(s.elements)-1].Max
}

func (s *StackWithMax) Pop() int {
	var element ElementWithCachedMax
	s.elements, element = s.elements[:len(s.elements)-1], s.elements[len(s.elements)-1]
	return element.Element
}

func (s *StackWithMax) Push(val int) {
	var newMax int
	if !s.Empty() {
		newMax = max(s.Max(), val)
	} else {
		newMax = val
	}
	s.elements = append(s.elements, ElementWithCachedMax{Element: val, Max: newMax})
}
