package null

func Map[K comparable, V any]() *NullableMap[K, V] {
	return new(NullableMap[K, V])
}

type NullableMap[K comparable, V any] struct {
	m map[K]V
}

func (m *NullableMap[K, V]) IsNull() bool {
	return m.m == nil
}

func (m *NullableMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *NullableMap[K, V]) Size() int {
	return len(m.m)
}

func (m *NullableMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *NullableMap[K, V]) Update(key K, update func(V) V, ifAbsent func() V) {
	if m.IsNull() {
		m.m = make(map[K]V)
	}

	if v, ok := m.m[key]; ok {
		m.m[key] = update(v)
	} else {
		m.m[key] = ifAbsent()
	}
}

func (m *NullableMap[K, V]) Set(key K, value V) {
	if m.IsNull() {
		m.m = make(map[K]V)
	}
	m.m[key] = value
}

func (m *NullableMap[K, V]) Get(key K) Nullable[V] {
	if v, ok := m.m[key]; ok {
		return Value(v)
	}
	return New[V]()
}

func (m *NullableMap[K, V]) Range() map[K]V {
	return m.m
}
