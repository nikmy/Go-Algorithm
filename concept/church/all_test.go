package church

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func evalInt(i Numeral) int {
	var (
		depth   = 0
		counter = func(Numeral) Numeral { depth++; return nil }
	)
	i(counter)(nil)
	return depth
}

func evalBool(b Boolean) bool {
	var (
		value bool
		mark  Term = func(Term) Term { value = true; return nil }
	)
	b(mark)(Zero)(nil)
	return value
}

func assertInt(t *testing.T, exp int, got Numeral, msgAndArgs ...any) {
	assert.Equal(t, exp, evalInt(got), msgAndArgs...)
}

func assertBool(t *testing.T, exp bool, got Boolean, msgAndArgs ...any) {
	assert.Equal(t, exp, evalBool(got), msgAndArgs...)
}

func assertFalse(t *testing.T, b Boolean) {
	assertBool(t, false, b)
}

func assertTrue(t *testing.T, b Boolean) {
	assertBool(t, true, b)
}

var nums []Numeral

func TestMain(m *testing.M) {
	nums = []Numeral{
		func(f Numeral) Numeral { return func(x Numeral) Numeral { return x } },
		func(f Numeral) Numeral { return func(x Numeral) Numeral { return f(x) } },
		func(f Numeral) Numeral { return func(x Numeral) Numeral { return f(f(x)) } },
	}

	for i := 3; i < 100; i++ {
		nums = append(nums, Inc(nums[len(nums)-1]))
	}

	os.Exit(m.Run())
}
