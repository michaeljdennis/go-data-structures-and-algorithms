package queue

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	n, ok := q.Dequeue()
	fmt.Println(n, ok)
}

func TestDequeue_QueueEmpty(t *testing.T) {
	q := NewQueue()
	n, ok := q.Dequeue()

	if ok {
		t.Errorf("queue should be empty - got: %+v", n)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	want := 2
	n, ok := q.Peek()

	if !ok {
		t.Fatal("queue should not be empty")
	}

	if want != n {
		t.Errorf("queue should have value %+v at the front", want)
	}
}

func TestPeek_QueueEmpty(t *testing.T) {
	q := NewQueue()

	got, ok := q.Peek()

	if ok {
		t.Errorf("queue should be empty - got %+v", got)
	}
}

func TestLen(t *testing.T) {
	q := NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	want := 3
	got := q.Len()

	if want != got {
		t.Errorf("queue length is incorrect - want: %+v, got: %+v", want, got)
	}
}
