package stack

// Stack ...
type Stack struct {
	top *Node
	len int
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{}
}

// Node ...
type Node struct {
	value int
	next  *Node
}

// NewNode ...
func NewNode(value int, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

// Push ...
func (s *Stack) Push(n int) {
	tmp := NewNode(n, s.top)
	s.top = tmp
	s.len++
}

// Pop ...
func (s *Stack) Pop() (int, bool) {
	if s.Len() == 0 {
		return 0, false
	}
	tmp := s.top.value
	if s.Len() > 1 {
		s.top = s.top.next
	}
	s.len--
	return tmp, true
}

// Peek ...
func (s *Stack) Peek() (int, bool) {
	if s.Len() == 0 {
		return 0, false
	}
	return s.top.value, true
}

// Empty ...
func (s *Stack) Empty() {
	s.top = nil
}

// Len ...
func (s *Stack) Len() int {
	return s.len
}
