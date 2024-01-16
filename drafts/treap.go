package drafts

import (
	"golang.org/x/exp/rand"
	"math"
)

type Treap struct {
	value int
	prior int

	max int

	left, right *Treap
}

func (t *Treap) Find(key int) bool {
	if t == nil {
		return false
	}

	if t.value == key {
		return true
	}

	if t.value < key {
		return t.right.Find(key)
	}

	return t.left.Find(key)
}

func (t *Treap) Insert(key int) *Treap {
	l, r := split(t, key)
	if l.getMax() == key {
		return merge(l, r)
	}

	return merge(merge(l, newTreapNode(key)), r)
}

func (t *Treap) Remove(key int) *Treap {
	if t == nil {
		return t
	}

	if t.value == key {
		return merge(t.left, t.right)
	}

	if t.value < key {
		return merge(t.left.Remove(key), t.right)
	}

	return merge(t.left, t.right.Remove(key))
}

func (t *Treap) getMax() int {
	if t == nil || t.right == nil {
		return math.MinInt
	}
	return t.max
}

func (t *Treap) getPrior() int {
	if t == nil {
		return 0
	}
	return t.prior
}

func newTreapNode(value int) *Treap {
	return &Treap{
		value: value,
		prior: rand.Int(),
	}
}

func merge(l, r *Treap) *Treap {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	if l.prior > r.prior {
		l.right = merge(l.right, r)
		return l
	}

	r.left = merge(l, r.left)
	return r
}

func split(t *Treap, key int) (*Treap, *Treap) {
	if t == nil {
		return nil, nil
	}

	var l, r *Treap

	if t.value <= key {
		rl, rr := split(t.right, key)
		t.right, l, r = rl, t, rr
		return l, r
	}

	ll, lr := split(t.left, key)
	t.left, l, r = lr, ll, t
	return l, r
}
