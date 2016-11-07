package shuffle

import "testing"

type arr []int

func (t arr) Swap(a int, b int) {
	t[a], t[b] = t[b], t[a]
}

func TestShuffle(t *testing.T) {
	a := arr([]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	})
	Shuffle(a, 0, len(a))
	// TODO what to assert against?

	a = arr([]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	})
	Shuffle(a, 1, len(a)-2)
	// TODO what to assert against?
}
