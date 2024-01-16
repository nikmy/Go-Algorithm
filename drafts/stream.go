package drafts

type Stream[E any] interface {
	Elem(int) E
	Len() int

	Slice() []E

	Filter(func(E) bool)
	ForEach(callback func(int, E) bool)
}

func NewStream[E any, S ~[]E](slice S) Stream[E] {
	s := &stream[E]{
		slice: slice,
		callback: func(int, E) bool {
			return true
		},
	}
	s.forEach = func(index int, elem E) bool {
		return s.callback(index, elem)
	}
	return s
}

type stream[E any] struct {
	slice    []E
	forEach  func(index int, elem E) bool
	callback func(index int, elem E) bool
}

func (s *stream[E]) Elem(i int) E {
	if i >= s.Len() {
		panic("index out of range")
	}
	return s.slice[i]
}

func (s *stream[E]) Len() int {
	var length int
	s.callback = func(int, E) bool {
		length++
		return true
	}
	s.finalize()
	return length
}

func (s *stream[E]) Slice() []E {
	result := make([]E, 0, len(s.slice))
	s.callback = func(_ int, elem E) bool {
		result = append(result, elem)
		return true
	}
	s.finalize()
	return result
}

func (s *stream[E]) Filter(f func(E) bool) {
	s.forEach = func(index int, elem E) bool {
		if !f(elem) {
			return true
		}
		return s.forEach(index, elem)
	}
}

func (s *stream[E]) ForEach(callback func(int, E) bool) {
	old := s.callback
	s.callback = func(index int, elem E) bool {
		if !old(index, elem) {
			return false
		}
		return callback(index, elem)
	}
}

func (s *stream[E]) finalize() {
	result := make([]E, s.Len())
	for i, elem := range s.slice {
		needContinue := s.forEach(i, elem)
		if !needContinue {
			break
		}
	}
	s.slice = result
}
