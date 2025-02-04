package heap_test

import (
	"testing"

	"nikand.dev/go/heap"
)

func TestHeapObj(tb *testing.T) {
	h := heap.New[int](heap.Less)

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

func ExampleHeap() {
	type Element struct{ X int }

	h := heap.Make[*Element](func(a []*Element, i, j int) bool {
		return a[i].X < a[j].X
	})

	h.Push(&Element{X: 5})
	h.Push(&Element{X: 2})
	h.Push(&Element{X: 10})

	_ = h.Pop() // &Element{X: 2}

	h.Data[0].X = 20 // 5 -> 20
	h.Fix(0)

	h.Del(0) // 10

	_ = h.Len() // 1

	// Output:
}
