package algo

// Trie
/////////////////////////////////////////////////////////////////////////////////////////
type node struct {
    next   map[uint8]int
    leafId int
}

func newNode(leafId int) node {
    return node{make(map[uint8]int), leafId}
}

type Trie struct {
    tree []node
}

func NewTrie() *Trie {
    return &Trie{[]node{newNode(-1)}}
}

func (t *Trie) TryNext(n int, c uint8) bool {
    _, can := t.tree[n].next[c]
    return can
}

func (t *Trie) GoNext(n int, c uint8) int {
    child, _ := t.tree[n].next[c]
    return child
}

func (t *Trie) Insert(id int, s string) {
    n := 0
    for i := 0; i < len(s); i++ {
        if t.TryNext(n, s[i]) {
            n = t.GoNext(n, s[i])
        } else {
            t.tree[n].next[s[i]] = len(t.tree)
            if i != len(s)-1 {
                n, t.tree = len(t.tree), append(t.tree, newNode(-1))
            } else {
                t.tree = append(t.tree, newNode(id))
            }
        }
    }
}

func (t *Trie) Find(s string) int {
    n := 0
    for i := 0; i < len(s); i++ {
        if t.TryNext(n, s[i]) {
            n = t.GoNext(n, s[i])
        } else {
            return -1
        }
    }
    return t.tree[n].leafId
}

/////////////////////////////////////////////////////////////////////////////////////////
