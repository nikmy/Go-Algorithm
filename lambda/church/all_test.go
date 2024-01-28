package church

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func assertInt(t *testing.T, exp int, got Numeral, msgAndArgs ...any) {
	assert.Equal(t, exp, makeInt(got), msgAndArgs...)
}

func assertBool(t *testing.T, exp bool, got Boolean, msgAndArgs ...any) {
	assert.Equal(t, exp, makeBool(got), msgAndArgs...)
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
