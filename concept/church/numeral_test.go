package church

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumerals(t *testing.T) {
	for i, n := range nums[:3] {
		assert.Equal(t, i, evalInt(n))
	}

	assert.Equal(t, evalInt(nums[0]), evalInt(Zero))
	assert.Equal(t, evalInt(nums[1]), evalInt(One))
	assert.NotEqual(t, evalInt(One), evalInt(Zero))
}

func TestInc(t *testing.T) {
	for i := 3; i < len(nums); i++ {
		assertInt(t, i, nums[i])
	}
}

func TestAdd(t *testing.T) {
	for i, a := range nums {
		for j, b := range nums {
			assertInt(t, i+j, Add2(a, b))
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
	assertInt(t, 0, Dec(Zero))
	for i := 1; i < len(nums); i++ {
		assertInt(t, i-1, Dec(nums[i]))
	}
}

func TestSub(t *testing.T) {
	for i, a := range nums[:5] {
		for j, b := range nums[:i] {
			assertInt(t, max(0, i-j), Sub2(a, b))
		}
	}
}

func TestDiv(t *testing.T) {
	for i, a := range nums {
		for j := 1; j < i; j++ {
			assertInt(t, i/j, Div2(a, nums[j]))
		}
	}
}

func TestMod(t *testing.T) {
	for i, a := range nums {
		for j := 1; j < i; j++ {
			assertInt(t, i%j, Mod2(a, nums[j]))
		}
	}
}
