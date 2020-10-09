package queue

// Queue ...
type Queue struct {
	front *Node
	len   int
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{}
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

// Enqueue ...
func (q *Queue) Enqueue(n int) {
	q.len++

	if q.front == nil {
		q.front = NewNode(n, nil)
		return
	}

	tmp := q.front
	node := NewNode(n, tmp)
	q.front = node
}

// Dequeue ...
func (q *Queue) Dequeue() (int, bool) {
	if q.front == nil {
		return 0, false
	}

	q.len--

	front := q.front
	q.front = q.front.next
	return front.value, true
}

// Peek ...
func (q *Queue) Peek() (int, bool) {
	if q.front == nil {
		return 0, false
	}

	return q.front.value, true
}

// Len ...
func (q *Queue) Len() int {
	return q.len
}
