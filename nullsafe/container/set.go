package container

type Set[T comparable] interface {
	Contains(elem T) bool
	Add(elems ...T)
	Remove(elems ...T)
}

func NewHashSet[T comparable]() Set[T] {
	return make(hashSet[T])
}

type hashSet[T comparable] map[T]struct{}

func (s hashSet[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s hashSet[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = struct{}{}
	}
}

func (s hashSet[T]) Remove(elems ...T) {
	for _, elem := range elems {
		delete(s, elem)
	}
}
