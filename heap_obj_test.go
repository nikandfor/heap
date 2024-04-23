package heap

import (
	"testing"
)

func TestHeapObj(tb *testing.T) {
	h := New[int](Less)

	h.Init([]int{10, 4, 2, 5, 1, 7, 8, 3, 6, 9, 0})

	last := -1

	for h.Len() != 0 {
		e := h.Pop()

		if last > e {
			tb.Errorf("%v // error", e)
			continue
		}

		tb.Logf("%v", e)
	}
}
