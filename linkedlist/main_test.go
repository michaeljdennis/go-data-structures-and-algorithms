package linkedlist

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	list := NewList()
	list.AddFirst(0)
	list.AddFirst(1)

	want := 1
	got := list.GetFirst()

	if want != got {
		t.Errorf("First node value in list incorrect: want \"%d\", got \"%d\"", want, got)
	}
}

func TestAddLast(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)

	want := 1
	got := list.GetLast()

	if want != got {
		t.Errorf("Last node in list incorrect: want %d, got %d", want, got)
	}
}

func TestDelete_FirstNode(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)
	list.Delete(0)

	want := 2
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect: want %d, got %d", want, got)
	}
}

func TestDelete_LastNode(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)
	list.Delete(2)

	want := 2
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect: want %d, got %d", want, got)
	}
}
