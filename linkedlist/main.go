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
func (l *List) AddFirst(n int) {
	tmp := NewNode(n, l.head)
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
func (l *List) AddLast(n int) {
	tmp := NewNode(n, nil)
	l.len++

	if l.head == nil {
		l.head = tmp
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = tmp
}

// GetLast ...
func (l *List) GetLast() int {
	if l.head == nil {
		return 0
	}

	tmp := l.head

	for tmp.next != nil {
		tmp = tmp.next
	}

	return tmp.value
}

// Reverse ...
func (l *List) Reverse() {
	var prev *Node
	tmp := l.head
	next := tmp.next

	for next != nil {
		next = tmp.next
		tmp.next = prev
		prev = tmp
		tmp = next
	}

	l.head = prev
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

// GetList ...
func (l *List) GetList() []int {
	tmp := l.head
	ls := []int{}

	if l.head == nil {
		return ls
	}

	for tmp != nil {
		ls = append(ls, tmp.value)
		tmp = tmp.next
	}

	return ls
}

// PrintList ...
func (l *List) PrintList() {
	fmt.Println(l.GetList())
}
