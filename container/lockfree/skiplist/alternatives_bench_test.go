package skiplist

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
	"testing"
	"unsafe"
)

func Benchmark_BinarySearch(b *testing.B) {
	sizes := []int64{10, 100, 1_000, 10_000, 100_000, 1_000_000}
	parallel := []int{1, 4, 16, 32}
	for _, p := range parallel {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d_parallel_%d_elements", p, size), func(b *testing.B) {
				b.ReportAllocs()
				b.SetParallelism(16)

				list := generateInts(0, int(size), 1)

				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					i := rand.Int64() % size
					for pb.Next() {
						if _, ok := slices.BinarySearch(list, i); !ok {
							panic(fmt.Sprintf("must be found"))
						}
						i += 7
						if i >= size {
							i -= size
						}
					}
				})
			})
		}
	}
}

func BenchmarkLock(b *testing.B) {
	b.SetParallelism(32)

	var lock sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			lock.Unlock()
		}
	})
}

func BenchmarkSharedLock(b *testing.B) {
	b.SetParallelism(32)

	var lock sync.RWMutex
	b.RunParallel(func(pb *testing.PB) {
		seed := uint64(uintptr(unsafe.Pointer(pb)))
		mask := uint64(1<<32 - 1)
		rnd := rand.New(rand.NewPCG(seed&mask, seed^mask))
		for pb.Next() {
			if rnd.Int()%3 == 0 {
				lock.RLock()
				lock.RUnlock()
			} else {
				lock.Lock()
				lock.Unlock()
			}
		}
	})
}
