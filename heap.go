package heap

import "cmp"

func Init[T any](d []T, k int, less func(d []T, i, j int) bool) bool {
	n := len(d)
	up := false

	for i := n/k - 1; i >= 0; i-- {
		up = up || Down(d, k, less, i)
	}

	return up
}

func Push[T any](d []T, k int, less func(d []T, i, j int) bool, e T) []T {
	d = append(d, e)

	_ = Up(d, k, less, len(d)-1)

	return d
}

func Pop[T any](d []T, k int, less func(d []T, i, j int) bool) (T, []T) {
	r := d[0]
	n := len(d) - 1
	d[0] = d[n]
	d = d[:n]

	_ = Down(d, k, less, 0)

	return r, d
}

func Fix[T any](d []T, k int, less func(d []T, i, j int) bool, i int) bool {
	return Up(d, k, less, i) || Down(d, k, less, i)
}

func Up[T any](d []T, k int, less func(d []T, i, j int) bool, i int) bool {
	i0 := i

	for {
		p := (i - 1) / k

		if p == i || !less(d, i, p) {
			break
		}

		swap(d, p, i)

		i = p
	}

	return i0 != i
}

func Down[T any](d []T, k int, less func(d []T, i, j int) bool, i int) bool {
	n := len(d)
	i0 := i

	for {
		b := i*k + 1
		if b >= n {
			break
		}

		j := b

		for c := 1; c < k && b+c < n; c++ {
			if less(d, b+c, j) {
				j = b + c
			}
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
