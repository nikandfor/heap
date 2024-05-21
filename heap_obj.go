package heap

type Heap[T any] struct {
	Data []T

	s    StatelessSwap[T]
	Less LessFunc[T]
	Swap SwapFunc[T]
}

func New[T any](less LessFunc[T]) *Heap[T] {
	return &Heap[T]{Less: less}
}

func (h *Heap[T]) Init(d []T) bool {
	h.Data = d
	return h.s.Init(d, h.Less, h.Swap)
}

func (h *Heap[T]) Len() int { return len(h.Data) }

func (h *Heap[T]) Push(e T) {
	h.Data = h.s.Push(h.Data, e, h.Less, h.Swap)
}

func (h *Heap[T]) Pop() T {
	r, d := h.s.Pop(h.Data, h.Less, h.Swap)
	h.Data = d

	return r
}

func (h *Heap[T]) Del(i int) {
	h.Data = h.s.Del(h.Data, i, h.Less, h.Swap)
}

func (h *Heap[T]) Fix(i int) bool {
	return h.s.Up(h.Data, i, h.Less, h.Swap) || h.s.Down(h.Data, i, h.Less, h.Swap)
}

func (h *Heap[T]) Up(i int) bool {
	return h.s.Up(h.Data, i, h.Less, h.Swap)
}

func (h *Heap[T]) Down(i int) bool {
	return h.s.Down(h.Data, i, h.Less, h.Swap)
}
