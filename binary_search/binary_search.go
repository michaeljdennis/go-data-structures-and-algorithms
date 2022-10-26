package binary_search

func BinarySearch(l []int, n int) int {
	if len(l) == 0 {
		return -1
	}

	if len(l) == 1 && l[0] == n {
		return -1
	}
	if len(l) == 1 && l[0] != n {
		return -1
	}

	if n < l[0] || n > l[len(l)-1] {
		return -1
	}

	return search(l, n, 0, l[len(l)-1])
}

func search(l []int, needle, low, high int) int {
	mid := (low + high) / 2 // 5

	if needle == l[mid] {
		return mid
	}

	if needle < l[mid] {
		return search(l, needle, 0, mid)
	}

	return search(l, needle, mid+1, len(l)-1)
}
