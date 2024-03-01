package memory

import (
	"math/bits"
	"unsafe"
)

const (
	tiny   = 16
	small  = 64
	medium = 256
	large  = 1024
	extra  = 4096
)

var bitsToSize = [13]int{
	tiny, tiny, tiny, tiny,
	small, small, small,
	medium, medium,
	large, large,
	extra, extra,
}

func NewArena(heapSize int) *Arena {
	return &Arena{
		heap: make([]byte, heapSize),
	}
}

type Arena struct {
	free  [][]byte
	freeT []*[tiny]byte
	freeS []*[small]byte
	freeM []*[medium]byte
	freeL []*[large]byte
	freeE []*[extra]byte

	heap []byte
	pos  int
}

func (a *Arena) Allocate(n int) []byte {
	if n < extra {
		chunkSize := bitsToSize[bits.Len(uint(n))+1]
		data := a.getChunk(chunkSize)
		return unsafe.Slice(data, n)
	}

	if n%tiny != 0 {
		n = (n + tiny) / tiny * tiny
	}

	if a.pos+n > len(a.heap) {
		panic("heap overflow")
	}

	data := &a.heap[a.pos]
	a.pos += n
	return unsafe.Slice(data, n)
}

func (a *Arena) Free(p uintptr, n int) { // TODO: join smaller buffers
	for n > extra {
		a.freeE = append(a.freeE, (*[extra]byte)(unsafe.Pointer(p)))
		n -= extra
		p += extra
	}
	for n > large {
		a.freeL = append(a.freeL, (*[large]byte)(unsafe.Pointer(p)))
		p += large
	}
	for n > medium {
		a.freeM = append(a.freeM, (*[medium]byte)(unsafe.Pointer(p)))
		p += medium
	}
	for n > small {
		a.freeS = append(a.freeS, (*[small]byte)(unsafe.Pointer(p)))
		p += small
	}
	for n > 0 {
		a.freeT = append(a.freeT, (*[tiny]byte)(unsafe.Pointer(p)))
		p += tiny
	}
}

func (a *Arena) getChunk(size int) *byte { // TODO: split bigger buffers
	var p *byte
	switch size {
	case tiny:
		if len(a.freeT) == 0 {
			last := len(a.freeT) - 1
			for i := range a.freeT[last] {
				a.freeT[last][i] = 0
			}
			p = &(*a.freeT[last])[0]
			a.freeT = a.freeT[:last]
		}
	case small:
		if len(a.freeS) > 0 {
			last := len(a.freeS) - 1
			for i := range a.freeS[last] {
				a.freeS[last][i] = 0
			}
			p = &(*a.freeS[0])[0]
			a.freeS = a.freeS[:last]
		}
	case medium:
		if len(a.freeM) > 0 {
			last := len(a.freeM) - 1
			for i := range a.freeM[last] {
				a.freeM[last][i] = 0
			}
			p = &(*a.freeM[0])[0]
			a.freeM = a.freeM[:last]
		}
	case large:
		if len(a.freeL) > 0 {
			last := len(a.freeL) - 1
			for i := range a.freeL[last] {
				a.freeL[last][i] = 0
			}
			p = &(*a.freeL[0])[0]
			a.freeL = a.freeL[:last]
		}
	case extra:
		if len(a.freeE) > 0 {
			last := len(a.freeE) - 1
			for i := range a.freeT[last] {
				a.freeE[last][i] = 0
			}
			p = &(*a.freeE[0])[0]
			a.freeE = a.freeE[:last]
		}
	default:
		panic("bad chunk size")
	}

	if p == nil {
		p = &a.heap[a.pos]
		a.pos += size
	}

	return p
}
