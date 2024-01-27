package church

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Int = Numeral
)

func eval(i Int) int {
	depth := 0
	counter := func(_ Int) Int {
		depth++
		return nil
	}

	i(counter)(Zero(nil))
	return depth
}

var n []Int

func TestMain(m *testing.M) {
	n = []Int{
		func(f Int) Int { return func(x Int) Int { return x } },
		func(f Int) Int { return func(x Int) Int { return f(x) } },
		func(f Int) Int { return func(x Int) Int { return f(f(x)) } },
	}

	for i := 3; i < 100; i++ {
		n = append(n, Inc(n[len(n)-1]))
	}

	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	for i := range n[:3] {
		assert.Equal(t, i, eval(n[i]))
	}

	assert.Equal(t, eval(n[0]), eval(Zero))
	assert.Equal(t, eval(n[1]), eval(One))
	assert.NotEqual(t, eval(One), eval(Zero))
}

func TestIncrement(t *testing.T) {
	for i := range n[3:] {
		assert.Equal(t, i, eval(n[i]))
	}
}

func TestAdd(t *testing.T) {
	for i, a := range n {
		for j, b := range n {
			assert.Equal(t, i+j, eval(Add2(a, b)))
		}
	}

}

func TestMul(t *testing.T) {
	for i, a := range n {
		for j, b := range n {
			exp, got := i*j, eval(Mul2(a, b))
			assert.Equal(t, exp, got, "%d * %d = %d", i, j, got)
		}
	}
}
