package heap

import "cmp"

type LessFunc[T any] func(d []T, i, j int) bool

func Init[T any](d []T, less LessFunc[T]) bool {
	n := len(d)
	up := false

	for i := n/2 - 1; i >= 0; i-- {
		up = up || Down(d, less, i)
	}

	return up
}

func PushTo[T any](d *[]T, less LessFunc[T], e T) {
	*d = Push(*d, less, e)
}

func Push[T any](d []T, less LessFunc[T], e T) []T {
	d = append(d, e)

	_ = Up(d, less, len(d)-1)

	return d
}

func PopFrom[T any](d *[]T, less LessFunc[T]) T {
	r, q := Pop(*d, less)
	*d = q

	return r
}

func Pop[T any](d []T, less LessFunc[T]) (T, []T) {
	r := d[0]
	n := len(d) - 1
	d[0] = d[n]
	d = d[:n]

	_ = Down(d, less, 0)

	return r, d
}

func Fix[T any](d []T, less LessFunc[T], i int) bool {
	return Up(d, less, i) || Down(d, less, i)
}

func Up[T any](d []T, less LessFunc[T], i int) bool {
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

func Down[T any](d []T, less LessFunc[T], i int) bool {
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

func swap[T any](d []T, i, j int) {
	d[i], d[j] = d[j], d[i]
}

func Less[T cmp.Ordered](d []T, i, j int) bool    { return d[i] < d[j] }
func Greater[T cmp.Ordered](d []T, i, j int) bool { return d[i] > d[j] }
