package container

import "github.com/nikmy/algo/nullsafe/null"

type Map[K comparable, V any] interface {
	IsEmpty() bool
	Size() int
	ContainsKey(key K) bool
	Set(key K, value V)
	Get(key K) null.Nullable[V]
	Range() map[K]V
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return make(hashMap[K, V])
}

type hashMap[K comparable, V any] map[K]V

func (m hashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m hashMap[K, V]) Size() int {
	return len(m)
}

func (m hashMap[K, V]) ContainsKey(key K) bool {
	_, ok := m[key]
	return ok
}

func (m hashMap[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m hashMap[K, V]) Get(key K) null.Nullable[V] {
	if v, ok := m[key]; ok {
		return null.Value(v)
	}
	return null.Null[V]()
}

func (m hashMap[K, V]) Range() map[K]V {
	return m
}
