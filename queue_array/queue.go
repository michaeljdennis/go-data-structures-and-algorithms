package queue

// Queue ...
type Queue struct {
	items []int
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue ...
func (q *Queue) Enqueue(n int) {
	q.items = append([]int{n}, q.items...)
}

// Dequeue ...
func (q *Queue) Dequeue() (int, bool) {
	if q.Len() == 0 {
		return 0, false
	}
	n, _ := q.Peek()
	q.items = q.items[0 : q.Len()-1]
	return n, true
}

// Peek ...
func (q *Queue) Peek() (int, bool) {
	if q.Len() == 0 {
		return 0, false
	}
	return q.items[q.Len()-1], true
}

// Len ...
func (q *Queue) Len() int {
	return len(q.items)
}
