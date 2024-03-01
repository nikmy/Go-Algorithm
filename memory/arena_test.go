package memory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArena(t *testing.T) {
	a := NewArena(1024)
	check(t, a)
	require.Equal(t, 256+16, a.pos)
}

func check(t *testing.T, a Allocator) {
	i := New[int](a)
	*i.Elem() = 9
	defer Delete(a, i)
	require.Equal(t, 9, *i.Elem())

	s := NewArray[byte](a, 200)
	defer DeleteArray(a, s)
	require.Len(t, 200, s)
}
