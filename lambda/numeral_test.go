package lambda

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Int = Numeral
)

func eval(i Int) int {
	var (
		depth = 0

		nop     Int = func(_ Int) Int { return nil }
		counter Int = func(_ Int) Int { depth++; return nil }
	)

	i(counter)(nop)
	return depth
}

func assertEquals(t *testing.T, want, got Int) {
	wantValue, gotValue := eval(want), eval(got)
	assert.Equal(t, wantValue, gotValue)
}

func assertNotEquals(t *testing.T, want, got Int) {
	wantValue, gotValue := eval(want), eval(got)
	assert.NotEqual(t, wantValue, gotValue)
}

var n []Int

func TestMain(m *testing.M) {
	n = []Int{
		func(f Int) Int { return func(x Int) Int { return x } },
		func(f Int) Int { return func(x Int) Int { return f(x) } },
		func(f Int) Int { return func(x Int) Int { return f(f(x)) } },
	}

	for i := 3; i < 35; i++ {
		n = append(n, Inc(n[len(n)-1]))
	}

	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	for i := range n {
		assert.Equal(t, i, eval(n[i]))
	}

	assertEquals(t, n[0], Zero)
	assertEquals(t, n[1], One)
	assertNotEquals(t, One, Zero)
}

func TestIncrement(t *testing.T) {
	assertEquals(t, n[1], Inc(n[0]))
	assertEquals(t, n[2], Inc(Inc(n[0])))
	assertEquals(t, Inc(n[2]), Inc(Inc(n[1])))
}

func TestAdd(t *testing.T) {
	for i := range n {
		assertEquals(t, n[i], Add2(n[i], Zero)) // Zero is neutral element
	}

	for i, a := range n {
		for j, b := range n {
			assert.Equal(t, i+j, eval(Add2(a, b)))  // numeric value
			assertEquals(t, Add2(a, b), Add2(b, a)) // commutativity

			for _, c := range n {
				assertEquals(t, Add2(a, Add2(b, c)), Add2(Add2(a, b), c)) // associativity
			}
		}
	}

}

func TestMul(t *testing.T) {
	for _, x := range n {
		assertEquals(t, Zero, Mul2(x, Zero)) // 0 * x = 0
		assertEquals(t, x, Mul2(x, One))     // 1 * x = x
	}

	for i, a := range n {
		for j, b := range n {
			assert.Equal(t, i*j, eval(Mul2(a, b)))  // numeric value
			assertEquals(t, Mul2(a, b), Mul2(b, a)) // commutativity

			for _, c := range n {
				assertEquals(t, Mul2(a, Mul2(b, c)), Mul2(Mul2(a, b), c))          // associativity
				assertEquals(t, Mul2(a, Add2(b, c)), Add2(Mul2(a, b), Add2(a, c))) // distributivity
			}
		}
	}
}
