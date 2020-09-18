package linkedlist

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	list := NewList()
	list.AddFirst(0)
	list.AddFirst(2)
	list.AddFirst(1)

	want := 1
	got, err := list.GetFirst()

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	if want != got {
		t.Errorf("First node value in list incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetFirst(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)

	want := 0
	got, err := list.GetFirst()

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	if want != got {
		t.Errorf("First node value in list incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetFirst_ListEmpty(t *testing.T) {
	list := NewList()

	got, err := list.GetFirst()

	if err == nil {
		t.Errorf("List should be empty and return an error - got: %d", got)
	}
}

func TestAddLast(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)

	want := 1
	got, err := list.GetLast()

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	if want != got {
		t.Errorf("Last node in list incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetLast(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)

	want := 2
	got, err := list.GetLast()

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	if want != got {
		t.Errorf("Last node value in list incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetLast_ListEmpty(t *testing.T) {
	list := NewList()

	got, err := list.GetLast()

	if err == nil {
		t.Errorf("List should be empty and return an error - got: %d", got)
	}
}

func TestReverse(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)
	list.Reverse()

	want := []int{2, 1, 0}
	got, err := list.GetList()

	if err != nil {
		t.Fatalf("error: %+v", err)
	}

	for i, v := range want {
		if v != got[i] {
			t.Errorf("Reversed list incorrect - want: %+v, got: %+v", want, got)
			break
		}
	}
}

func TestDelete_FirstNode(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)
	err := list.Delete(0)

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	want := 2
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect - want: %d, got: %d", want, got)
	}
}

func TestDelete_LastNode(t *testing.T) {
	list := NewList()
	list.AddLast(0)
	list.AddLast(1)
	list.AddLast(2)
	err := list.Delete(2)

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	want := 2
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect - want: %d, got: %d", want, got)
	}
}

func TestDelete_ListEmpty(t *testing.T) {
	list := NewList()

	list.Delete(0)
}

func TestLen(t *testing.T) {
	list := NewList()
	list.AddFirst(0)
	list.AddFirst(1)

	want := 2
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect - want: %d, got: %d", want, got)
	}
}

func TestLen_ListEmpty(t *testing.T) {
	list := NewList()

	want := 0
	got := list.Len()

	if want != got {
		t.Errorf("List length incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetList(t *testing.T) {
	list := NewList()
	list.AddFirst(0)
	list.AddFirst(1)
	list.AddFirst(2)

	sl, err := list.GetList()

	if err != nil {
		t.Errorf("error: %+v", err)
	}

	want := 3
	got := len(sl)

	if want != got {
		t.Errorf("List length incorrect - want: %d, got: %d", want, got)
	}
}

func TestGetList_ListEmpty(t *testing.T) {
	list := NewList()

	sl, err := list.GetList()

	want := 0
	got := list.Len()

	if want != got || err == nil {
		t.Errorf("List length incorrect - want: %d, got: %d", 0, len(sl))
	}
}

func TestPrintList(t *testing.T) {
	list := NewList()
	list.AddFirst(0)
	list.PrintList()
}
