package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	want := 3
	got := stack.Len()

	if want != got {
		t.Errorf("stack length incorrect - want: %d, got: %d", want, got)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	got, ok := stack.Pop()
	if !ok {
		t.Error("no value popped")
	}

	want := 3
	if want != stack.Len() {
		t.Errorf("stack length incorrect - want: %d, got: %d", want, stack.Len())
	}

	want = 3
	if want != got {
		t.Errorf("wrong value popped - want: %d, got: %d", want, got)
	}
}

func TestPop_StackEmpty(t *testing.T) {
	stack := NewStack()

	got, ok := stack.Pop()
	if ok {
		t.Errorf("stack should be empty - got: %+v", got)
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	got, ok := stack.Peek()
	if !ok {
		t.Error("no value popped")
	}

	want := 5
	if want != stack.Len() {
		t.Errorf("stack length incorrect - want: %d, got: %d", want, got)
	}

	want = 4
	if want != got {
		t.Errorf("wrong value peeked - want: %d, got: %d", want, got)
	}
}

func TestPeek_StackEmpty(t *testing.T) {
	stack := NewStack()

	got, ok := stack.Peek()
	if ok {
		t.Errorf("stack should be empty - got: %+v", got)
	}
}

func TestEmpty(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	want := 3
	got := stack.Len()
	if want != got {
		t.Errorf("stack length incorrect - want: %d, got: %d", want, got)
	}

	stack.Empty()

	want = 0
	got = stack.Len()
	if want != got {
		t.Errorf("stack length incorrect - want: %d, got: %d", want, got)
	}
}
