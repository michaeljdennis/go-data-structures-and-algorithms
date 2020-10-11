package stack

// Stack ...
type Stack struct {
	items []int
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{}
}

// Push ...
func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

// Pop ...
func (s *Stack) Pop() (int, bool) {
	if s.Len() == 0 {
		return 0, false
	}
	tmp := s.items[s.Len()-1]
	s.items = s.items[0 : s.Len()-1]
	return tmp, true
}

// Peek ...
func (s *Stack) Peek() (int, bool) {
	if s.Len() == 0 {
		return 0, false
	}
	return s.items[s.Len()-1], true
}

// Empty ...
func (s *Stack) Empty() {
	s.items = make([]int, 0)
}

// Len ...
func (s *Stack) Len() int {
	return len(s.items)
}
