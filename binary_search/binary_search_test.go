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

func TestBinarySearch_ValueNotInList(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15}
	expectedValue := 13
	expectedIndex := -1
	actualIndex := BinarySearch(l, expectedValue)

	if actualIndex != expectedIndex {
		t.Errorf("expected %d at index %d in the list %+v", expectedValue, expectedIndex, l)
	}
}

func TestBinarySearch_EmptyList(t *testing.T) {
	l := []int{}
	expectedValue := 13
	expectedIndex := -1
	actualIndex := BinarySearch(l, expectedValue)

	if actualIndex != expectedIndex {
		t.Errorf("expected %d at index %d in the list %+v", expectedValue, expectedIndex, l)
	}
}
