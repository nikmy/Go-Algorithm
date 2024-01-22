package hyper

import (
	"github.com/nikmy/algo/container/bitset"
)

// HittingSet returns list of vertices that forms hitting set
func (h *Hypergraph) HittingSet() []int {
	hs := make([]int, 0)
	uncovered := bitset.New(len(h.e2v))
	uncovered.Flip()

	rest := make([]*bitset.Bitset, len(h.v2e))
	copy(rest, h.v2e)

	for uncovered.Any() {
		i, next, cnt := 0, rest[0], rest[0].Count()
		for j := 1; j < len(rest); j++ {
			newNext := bitset.Xor(rest[j], uncovered)
			newCount := newNext.Count()
			if newCount > cnt {
				i, next, cnt = j, newNext, newCount
			}
		}

		hs = append(hs, i)

		uncovered = bitset.Xor(uncovered, next)
		rest[i], rest[len(rest)-1] = rest[len(rest)-1], rest[i]
		rest = rest[:len(rest)-1]
	}
	return hs
}
