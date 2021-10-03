package binarytree

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	bt := NewBinaryTree()

	if bt.root != nil {
		t.Error("root node should be nil")
	}
}

func TestInsert(t *testing.T) {
	b := NewBinaryTree()
	b.Insert(10)
	b.Insert(8)
	b.Insert(13)
	b.Insert(5)
	b.Insert(9)
	b.Insert(12)
	b.Insert(25)

	want := []int{5, 8, 9, 10, 12, 13, 25}
	var got []int

	b.InOrderTraversal(b.root, func(n int) {
		got = append(got, n)
	})

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("node mismatch - want: %d, got %d", want[i], got[i])
		}
	}
}
