package std

type Hashable[K comparable] interface {
	Hash() K
}

func Hash[T any, H comparable](x T) any {
	switch h := any(x).(type) {
	case H:
		return h
	case Hashable[H]:
		return h.Hash()
	default:
		return any(x)
	}
}

type mapEntry[K, V any] struct {
	k K
	v V
}

func NewHashMap[H comparable, K Hashable[H], V any]() HashMap[H, K, V] {
	return make(HashMap[H, K, V])
}

type HashMap[H comparable, K Hashable[H], V any] map[H]mapEntry[K, V]

func (m HashMap[H, K, V]) Len() int {
	return len(m)
}

func (m HashMap[H, K, V]) Set(key K, value V) {
	m[key.Hash()] = mapEntry[K, V]{k: key, v: value}
}

func (m HashMap[H, K, V]) Delete(key K) {
	delete(m, key.Hash())
}

func (m HashMap[H, K, V]) Get(key K) V {
	return m[key.Hash()].v
}

func (m HashMap[H, K, V]) Has(key K) bool {
	_, ok := m[key.Hash()]
	return ok
}

func (m HashMap[H, K, V]) HasGet(key K) (V, bool) {
	entry, ok := m[key.Hash()]
	return entry.v, ok
}

func (m HashMap[H, K, V]) ForEach(yield func(k K, v V)) {
	for _, entry := range m {
		yield(entry.k, entry.v)
	}
}

func NewHashSet[H comparable, E Hashable[H]]() HashSet[H, E] {
	return make(HashSet[H, E])
}

type HashSet[H comparable, E Hashable[H]] map[H]E

func (s HashSet[H, E]) Len() int {
	return len(s)
}

func (s HashSet[H, E]) Has(x E) bool {
	_, ok := s[x.Hash()]
	return ok
}

func (s HashSet[H, E]) Add(elems ...E) {
	for _, elem := range elems {
		s[elem.Hash()] = elem
	}
}

func (s HashSet[H, E]) Delete(elems ...E) {
	for _, elem := range elems {
		delete(s, elem.Hash())
	}
}

func (s HashSet[H, E]) ForEach(yield func(elem E)) {
	for _, elem := range s {
		yield(elem)
	}
}

func (s HashSet[H, E]) ToSlice() []E {
	slice := make([]E, 0, s.Len())
	s.ForEach(func(elem E) { slice = append(slice, elem) })
	return slice
}
