package church

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func evalBool(b Boolean) bool {
	var (
		value  bool
		mark   Term = func(Term) Term { value = true; return nil }
		unmark Term = func(Term) Term { value = false; return nil }
	)

	b(mark)(unmark)(False)

	return value
}

func assertFalse(t *testing.T, b Boolean) {
	assert.False(t, evalBool(b))
}

func assertTrue(t *testing.T, b Boolean) {
	assert.True(t, evalBool(b))
}

func TestBool(t *testing.T) {
	assertFalse(t, False)
	assertTrue(t, True)
}

func TestNot(t *testing.T) {
	assertFalse(t, Not(True))
	assertTrue(t, Not(False))
}

func TestAnd(t *testing.T) {
	assertFalse(t, And2(False, False))
	assertFalse(t, And2(False, True))
	assertFalse(t, And2(True, False))
	assertTrue(t, And2(True, True))
}

func TestOr(t *testing.T) {
	assertFalse(t, Or2(False, False))
	assertTrue(t, Or2(False, True))
	assertTrue(t, Or2(True, False))
	assertTrue(t, Or2(True, True))
}

func TestXor(t *testing.T) {
	assertFalse(t, Xor2(False, False))
	assertTrue(t, Xor2(True, False))
	assertTrue(t, Xor2(False, True))
	assertFalse(t, Xor2(True, True))
}

func TestPair(t *testing.T) {
	p := Pair2(True, False)
	assertTrue(t, Left(p))
	assertFalse(t, Right(p))

	assert.Equal(t, 0, evalInt(Left(Pair2(Zero, One))))
	assert.Equal(t, 1, evalInt(Right(Pair2(Zero, One))))
}
