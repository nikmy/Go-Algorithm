package church

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func evalInt(i Numeral) int {
	depth := 0
	counter := func(_ Numeral) Numeral {
		depth++
		return nil
	}

	i(counter)(Zero(nil))
	return depth
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

func TestNumerals(t *testing.T) {
	for i, n := range nums[:3] {
		assert.Equal(t, i, evalInt(n))
	}

	assert.Equal(t, evalInt(nums[0]), evalInt(Zero))
	assert.Equal(t, evalInt(nums[1]), evalInt(One))
	assert.NotEqual(t, evalInt(One), evalInt(Zero))
}

func TestIncrement(t *testing.T) {
	for i, n := range nums[3:] {
		assert.Equal(t, i, evalInt(n))
	}
}

func TestAdd(t *testing.T) {
	for i, a := range nums {
		for j, b := range nums {
			assert.Equal(t, i+j, evalInt(Add2(a, b)))
		}
	}

}

func TestMul(t *testing.T) {
	for i, a := range nums {
		for j, b := range nums {
			exp, got := i*j, evalInt(Mul2(a, b))
			assert.Equal(t, exp, got, "%d * %d = %d", i, j, got)
		}
	}
}

func TestDec(t *testing.T) {
	for i := 1; i < len(nums); i++ {
		assert.Equal(t, i-1, evalInt(Dec(nums[i])))
	}
	assert.Equal(t, 0, evalInt(Dec(Zero)))
}

func TestSub(t *testing.T) {
	for i, a := range nums {
		for j, b := range nums {
			assert.Equal(t, max(0, i-j), evalInt(Sub2(a, b)))
		}
	}
}
