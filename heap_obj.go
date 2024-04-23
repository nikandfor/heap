package heap

type Heap[T any] struct {
	Data []T
	K    int
	Less func(d []T, i, j int) bool
}

func New[T any](k int, less func(s []T, i, j int) bool) *Heap[T] {
	return &Heap[T]{K: k, Less: less}
}

func (h *Heap[T]) Init(d []T) bool {
	h.Data = d
	return Init(d, h.K, h.Less)
}

func (h *Heap[T]) Len() int { return len(h.Data) }

func (h *Heap[T]) Push(e T) {
	h.Data = Push(h.Data, h.K, h.Less, e)
}

func (h *Heap[T]) Pop() T {
	r, d := Pop(h.Data, h.K, h.Less)
	h.Data = d

	return r
}

func (h *Heap[T]) Fix(i int) bool {
	return h.Up(i) || h.Down(i)
}

func (h *Heap[T]) Up(i int) bool {
	return Up(h.Data, h.K, h.Less, i)
}

func (h *Heap[T]) Down(i int) bool {
	return Down(h.Data, h.K, h.Less, i)
}
