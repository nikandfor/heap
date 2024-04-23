package heap

import (
	"math/rand"
	"testing"
)

func BenchmarkHeapInit(tb *testing.B) {
	less := func(d []int, i, j int) bool { return d[i] < d[j] }

	q := make([]int, 1_000_000)
	d := make([]int, len(q))
	rnd := rand.New(rand.NewSource(0))

	for i := range q {
		q[i] = rnd.Int()
	}

	tb.ResetTimer()

	for i := 0; i < tb.N; i++ {
		copy(d, q)

		Init(d, less)
	}
}
