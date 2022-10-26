package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedValue := 8
	expectedIndex := 7
	actualIndex := BinarySearch(l, expectedValue)

	if actualIndex != expectedIndex {
		t.Errorf("expected %d at index %d in the list %+v", expectedValue, expectedIndex, l)
	}
}
