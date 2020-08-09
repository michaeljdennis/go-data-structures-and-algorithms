package linkedlist

import "fmt"

// List ...
type List struct {
	head *Node
	len  int
}

// NewList ...
func NewList() *List {
	return &List{}
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

// AddFirst ...
func (l *List) AddFirst(s int) {
	tmp := NewNode(s, l.head)
	l.head = tmp
	l.len++
}

// GetFirst ...
func (l *List) GetFirst() int {
	if l.head == nil {
		return 0
	}
	return l.head.value
}

// AddLast ...
func (l *List) AddLast(s int) {
	current := l.head
	tmp := NewNode(s, nil)

	for {
		if l.head == nil {
			l.head = tmp
			break
		}
		if current.next != nil {
			current = current.next
			continue
		}
		current.next = tmp
		break
	}

	l.len++
}

// GetLast ...
func (l *List) GetLast() int {
	if l.head == nil {
		return 0
	}

	tmp := l.head

	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}

	return tmp.value
}

// Reverse ...
func (l *List) Reverse() {
	// TODO:
}

// Delete ...
func (l *List) Delete(n int) {
	if l.head == nil {
		return
	}

	var prev *Node
	tmp := l.head

	for {
		if tmp.value == n {
			// If first, just set head to the next node
			if prev == nil {
				l.head = tmp.next
			} else {
				prev.next = tmp.next
			}
			l.len--
			break
		}

		prev = tmp
		tmp = tmp.next
	}
}

// Len ...
func (l *List) Len() int {
	return l.len
}

// PrintList ...
func (l *List) PrintList() {
	current := l.head
	s := []int{}
	var done bool

	if l.head == nil {
		fmt.Println(s)
		return
	}

	for !done {
		s = append(s, current.value)
		if current.next == nil {
			done = true
		}
		current = current.next
	}
	fmt.Println(s)
}
