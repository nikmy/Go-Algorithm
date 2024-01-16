package strings

func NewTrie(words ...string) *Trie {
	d := &Trie{[]node{newNode(false)}}
	for _, word := range words {
		d.Insert(word)
	}
	return d
}

type Trie struct {
	tree []node
}

func (t *Trie) Insert(s string) {
	n := 0
	for i, c := range []byte(s) {
		if t.can(n, c) {
			n = t.goBy(n, c)
			continue
		}

		t.tree[n].next[c-'a'+1] = len(t.tree)
		if i != len(s)-1 {
			n, t.tree = len(t.tree), append(t.tree, newNode(false))
		} else {
			t.tree = append(t.tree, newNode(true))
		}
	}
}

func (t *Trie) Delete(s string) {
	n := 0
	for _, c := range []byte(s) {
		if !t.can(n, c) {
			return
		}
		n = t.goBy(n, c)
	}
	t.tree[n].isLeaf = false
}

func (t *Trie) Has(s string) bool {
	n := 0
	for i := 0; i < len(s); i++ {
		if !t.can(n, s[i]) {
			return false
		}
		n = t.goBy(n, s[i])
	}
	return t.tree[n].isLeaf
}

func (t *Trie) can(n int, c byte) bool {
	return t.tree[n].next[c] != 0
}

func (t *Trie) goBy(n int, c byte) int {
	return t.tree[n].next[c]
}

func newNode(leaf bool) node {
	return node{isLeaf: leaf}
}

type node struct {
	next   [256]int
	isLeaf bool
}
