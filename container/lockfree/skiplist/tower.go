package skiplist

import (
	"math/rand/v2"
	"sync/atomic"
)

const (
	maxLevel = 32
)

func newTower(x int64) *tower {
	levels := 1
	for levels < maxLevel && rand.Int()%4 == 0 {
		levels++
	}

	return &tower{
		elem: x,
		next: make([]atomic.Pointer[tower], levels),
	}
}

type tower struct {
	elem int64
	next []atomic.Pointer[tower]

	state atomic.Int32
}

func (t *tower) find(x int64) bool {
	node := t

	for level := len(t.next) - 1; level >= 0; level-- {
		for node != nil && node.elem < x {
			next := node.next[level].Load()
			if next == nil || next.elem > x {
				break
			}
			node = next
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

func (t *tower) findLinks(links []*tower, x int64) (int, *tower) {
	var (
		node = t
		next *tower
	)
	for level := len(t.next) - 1; level >= 0; level-- {
		next = node.next[level].Load()
		for next != nil && next.elem < x {
			node = next
			next = next.next[level].Load()
		}
		links[level] = node
	}

	if next != nil && next.elem == x {
		return len(t.next), next
	}
	return len(t.next), nil
}

// towerState represents state of a tower:
//  1. tower is created in INIT state;
//  2. when linked, it turns into CREATED state;
//  3. when tower.unlink is called in CREATED state, it turns into DELETING state
type towerState = int32

const (
	towerStateInit towerState = iota
	towerStateCreated
	towerStateDeleting
)

func (t *tower) link(links []*tower) bool {
	for level := 0; level < len(t.next); level++ {
		left := links[level]
		for {
			right := left.next[level].Load()
			for right != nil && right.elem < t.elem {
				left = right
				right = right.next[level].Load()
			}
			if level == 0 && right != nil && right.elem == t.elem {
				return false
			}

			t.next[level].Store(right)
			if left.next[level].CompareAndSwap(right, t) {
				break
			}
		}
	}

	t.state.Store(towerStateCreated)

	return true
}

func (t *tower) unlink(links []*tower) bool {
	if !t.state.CompareAndSwap(towerStateCreated, towerStateDeleting) {
		return false
	}

	for level := 0; level < len(t.next); level++ {
		/*
			Unlinking B from A -> B -> C

			1. Make loop:

				A -> (B)

			2. Switch forward link:

				A   (B)   C
			    |---------^

			3. Make reverse link:

				A <--- B    C
				|-----------^
		*/

		// Step 1: announce deletion (make a loop link)
		var right *tower
		for {
			right = t.next[level].Load()
			if t.next[level].CompareAndSwap(right, t) {
				break
			}
		}

		// Step 2: switch forward link
		left := links[level]
		for {
			next := left.next[level].Load()
			for next != nil && next.elem < t.elem {
				left = next
				next = left.next[level].Load()
			}
			if next == nil || next.elem > t.elem {
				break
			}
			if left.next[level].CompareAndSwap(t, right) {
				break
			}
		}

		// Step 3: make reverse link
		if t.next[level].CompareAndSwap(t, left) {
			if t.next[level].Load() != left {
				panic("tower unlink fail")
			}
		}
	}

	return true
}
