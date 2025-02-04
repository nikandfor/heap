package heap_test

import (
	"math/rand"
	"testing"
	"time"

	"nikand.dev/go/heap"
)

func BenchmarkHeapInit(tb *testing.B) {
	var h heap.Stateless[int]
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

		h.Init(d, less)
	}
}

type (
	Gossip struct {
		Peers []*Peer // Peers is a standalone field

		h heap.Stateless[*Peer] // used to sort peers by ping time
	}

	Peer struct {
		LastPing time.Time
	}
)

func (s *Gossip) less(ps []*Peer, i, j int) bool {
	return ps[i].LastPing.Before(ps[j].LastPing)
}

func ExampleStateless() {
	s := &Gossip{}

	// init some peers
	for i := 0; i < 5; i++ {
		s.Peers = append(s.Peers, &Peer{})
	}

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for now := range t.C {
		if len(s.Peers) == 0 {
			continue
		}

		p := s.Peers[0] // first in the queue

		// Ping
		p.LastPing = now
		s.h.Fix(s.Peers, 0, s.less) // move p to its new position in the queue
	}
}
