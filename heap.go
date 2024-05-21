package heap

import "cmp"

type (
	LessFunc[T any] func(d []T, i, j int) bool
	SwapFunc[T any] func(d []T, i, j int)

	Stateless[T any]     struct{}
	StatelessSwap[T any] struct{}
)

func (h Stateless[T]) Init(d []T, less LessFunc[T]) bool {
	return (StatelessSwap[T]{}).Init(d, less, nil)
}

func (h Stateless[T]) PushTo(d *[]T, e T, less LessFunc[T]) {
	(StatelessSwap[T]{}).PushTo(d, e, less, nil)
}

func (h Stateless[T]) Push(d []T, e T, less LessFunc[T]) []T {
	return (StatelessSwap[T]{}).Push(d, e, less, nil)
}

func (h Stateless[T]) PopFrom(d *[]T, less LessFunc[T]) T {
	return (StatelessSwap[T]{}).PopFrom(d, less, nil)
}

func (h Stateless[T]) Pop(d []T, less LessFunc[T]) (T, []T) {
	return (StatelessSwap[T]{}).Pop(d, less, nil)
}

func (h Stateless[T]) Del(d []T, i int, less LessFunc[T]) []T {
	return (StatelessSwap[T]{}).Del(d, i, less, nil)
}

func (h Stateless[T]) Fix(d []T, i int, less LessFunc[T]) bool {
	return (StatelessSwap[T]{}).Fix(d, i, less, nil)
}

func (h Stateless[T]) Up(d []T, i int, less LessFunc[T]) bool {
	return (StatelessSwap[T]{}).Up(d, i, less, nil)
}

func (h Stateless[T]) Down(d []T, i int, less LessFunc[T]) bool {
	return (StatelessSwap[T]{}).Down(d, i, less, nil)
}

func (h StatelessSwap[T]) Init(d []T, less LessFunc[T], swap SwapFunc[T]) bool {
	n := len(d)
	up := false

	for i := n/2 - 1; i >= 0; i-- {
		up = up || h.Down(d, i, less, swap)
	}

	return up
}

func (h StatelessSwap[T]) PushTo(d *[]T, e T, less LessFunc[T], swap SwapFunc[T]) {
	*d = h.Push(*d, e, less, swap)
}

func (h StatelessSwap[T]) Push(d []T, e T, less LessFunc[T], swap SwapFunc[T]) []T {
	d = append(d, e)

	_ = h.Up(d, len(d)-1, less, swap)

	return d
}

func (h StatelessSwap[T]) PopFrom(d *[]T, less LessFunc[T], swap SwapFunc[T]) T {
	r, q := h.Pop(*d, less, swap)
	*d = q

	return r
}

func (h StatelessSwap[T]) Pop(d []T, less LessFunc[T], swap SwapFunc[T]) (T, []T) {
	r := d[0]
	n := len(d) - 1
	d[0] = d[n]
	d = d[:n]

	_ = h.Down(d, 0, less, swap)

	return r, d
}

func (h StatelessSwap[T]) Del(d []T, i int, less LessFunc[T], swap SwapFunc[T]) []T {
	if swap == nil {
		swap = defswap
	}

	if i == len(d)-1 {
		return d[:len(d)-1]
	}

	swap(d, i, len(d)-1)

	d = d[:len(d)-1]

	h.Fix(d, i, less, swap)

	return d
}

func (h StatelessSwap[T]) Fix(d []T, i int, less LessFunc[T], swap SwapFunc[T]) bool {
	return h.Up(d, i, less, swap) || h.Down(d, i, less, swap)
}

func (h StatelessSwap[T]) Up(d []T, i int, less LessFunc[T], swap SwapFunc[T]) bool {
	if swap == nil {
		swap = defswap
	}

	i0 := i

	for {
		p := (i - 1) / 2

		if p == i || !less(d, i, p) {
			break
		}

		swap(d, p, i)

		i = p
	}

	return i0 != i
}

func (h StatelessSwap[T]) Down(d []T, i int, less LessFunc[T], swap SwapFunc[T]) bool {
	if swap == nil {
		swap = defswap
	}

	n := len(d)
	i0 := i

	for {
		b := i*2 + 1
		if b >= n {
			break
		}

		j := b

		if b+1 < n && less(d, b+1, j) {
			j = b + 1
		}

		if !less(d, j, i) {
			break
		}

		swap(d, j, i)
		i = j
	}

	return i0 != i
}

func defswap[T any](d []T, i, j int) {
	d[i], d[j] = d[j], d[i]
}

func Less[T cmp.Ordered](d []T, i, j int) bool    { return d[i] < d[j] }
func Greater[T cmp.Ordered](d []T, i, j int) bool { return d[i] > d[j] }
