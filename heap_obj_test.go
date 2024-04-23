package heap

import (
	"fmt"
	"testing"
)

func TestHeapObj(tb *testing.T) {
	testHeapObj(tb, 2)
	testHeapObj(tb, 3)
	testHeapObj(tb, 4)
	testHeapObj(tb, 6)
	testHeapObj(tb, 12)
}

func testHeapObj(tb *testing.T, k int) {
	tb.Helper()

	tb.Run(fmt.Sprintf("k%d", k), func(tb *testing.T) {
		h := New[int](2, Less)

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
	})
}
