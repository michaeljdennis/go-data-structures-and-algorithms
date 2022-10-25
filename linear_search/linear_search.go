package linear_search

func Search(l []int, n int) (int, int) {
	for i, v := range l {
		if v == n {
			return i, v
		}
	}
	return -1, 0
}
