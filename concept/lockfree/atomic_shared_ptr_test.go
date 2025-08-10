package concept

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestAtomicSharedPtr_size(t *testing.T) {
	require.Equal(t, uintptr(16), unsafe.Sizeof(AtomicSharedPtr[int]{}))
}


