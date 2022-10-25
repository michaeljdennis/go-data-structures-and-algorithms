package linear_search

import "testing"

func TestSearch(t *testing.T) {
	list := []int{8, 13, 16, 4, 23, 42, 15}
	index, value := Search(list, 13)

	want := 1
	if index != want {
		t.Errorf("wanted %d, got %d", want, index)
	}

	want = 13
	if value != want {
		t.Errorf("wanted %d, got %d", want, value)
	}
}
