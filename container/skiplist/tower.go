package skiplist

import (
	"math/rand/v2"
)

const (
	maxLevel = 32
)

func newTower(x int64) *tower {
	levels := 1
	for levels < maxLevel && rand.Int()%2 == 0 {
		levels++
	}

	return &tower{
		elem: x,
		next: make([]*tower, levels),
	}
}

type tower struct {
	elem int64
	next []*tower
}

func (t *tower) find(x int64) bool {
	node := t

	for level := len(t.next) - 1; level >= 0; level-- {
		for node != nil && node.next[level] != nil && node.next[level].elem <= x {
			node = node.next[level]
		}
		if node == nil {
			return false
		}
		if node.elem == x {
			return true
		}
	}

	return false
}

func (t *tower) findLinks(x int64) []*tower {
	links := make([]*tower, len(t.next))

	node := t
	for level := len(t.next) - 1; level >= 0; level-- {
		next := node.next[level]
		for next != nil && next.elem < x {
			node = next
			next = next.next[level]
		}
		links[level] = node
	}

	return links
}

func (t *tower) unlink(links []*tower) bool {
	for level := 0; level < len(t.next); level++ {
		right := t.next[level]
		links[level].next[level] = right
	}

	return true
}

func (t *tower) link(links []*tower) bool {
	for level := 0; level < len(t.next); level++ {
		link := links[level]

		right := link.next[level]
		t.next[level] = right
		link.next[level] = t
	}

	return true
}
