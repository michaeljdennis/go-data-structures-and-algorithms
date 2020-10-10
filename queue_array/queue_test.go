package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	want := 3
	got := q.Len()

	if want != got {
		t.Errorf("queue length incorrect - want: %d, got: %d", want, got)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	got, ok := q.Dequeue()

	if !ok {
		t.Error("no value dequeued")
	}

	want := 2
	if want != q.Len() {
		t.Errorf("queue length incorrect - want: %d, got: %d", want, q.Len())
	}

	want = 3
	if want != got {
		t.Errorf("wrong value dequeued - want: %d, got: %d", want, got)
	}
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

	want := 0
	got, ok := q.Peek()

	if !ok {
		t.Fatal("queue should not be empty")
	}

	if want != got {
		t.Errorf("peek value incorrect - want: %d, got: %d", want, got)
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
		t.Errorf("queue length incorrect - want: %+v, got: %+v", want, got)
	}
}
