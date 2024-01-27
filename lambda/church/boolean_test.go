package church

import "testing"

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

	assertInt(t, 0, Left(Pair2(Zero, One)))
	assertInt(t, 1, Right(Pair2(Zero, One)))
}
