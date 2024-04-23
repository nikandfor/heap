package heap

type Heap[T any] struct {
	Data []T
	Less LessFunc[T]
}

func New[T any](less LessFunc[T]) *Heap[T] {
	return &Heap[T]{Less: less}
}

func (h *Heap[T]) Init(d []T) bool {
	h.Data = d
	return Init(d, h.Less)
}

func (h *Heap[T]) Len() int { return len(h.Data) }

func (h *Heap[T]) Push(e T) {
	h.Data = Push(h.Data, h.Less, e)
}

func (h *Heap[T]) Pop() T {
	r, d := Pop(h.Data, h.Less)
	h.Data = d

	return r
}

func (h *Heap[T]) Fix(i int) bool {
	return Up(h.Data, h.Less, i) || Down(h.Data, h.Less, i)
}

func (h *Heap[T]) Up(i int) bool {
	return Up(h.Data, h.Less, i)
}

func (h *Heap[T]) Down(i int) bool {
	return Down(h.Data, h.Less, i)
}
