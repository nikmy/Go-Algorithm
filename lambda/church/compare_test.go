package church

import "testing"

func TestCompare(t *testing.T) {
	// IsZero
	assertTrue(t, IsZero(Zero))
	for _, n := range nums[1:] {
		assertFalse(t, IsZero(n))
	}

	// LessOrEqual
	for i, a := range nums {
		for j, b := range nums {
			assertBool(t, i <= j, LessOrEqual2(a, b))
		}
	}

	// Less
	for i, a := range nums {
		for j, b := range nums {
			assertBool(t, i < j, Less2(a, b))
		}
	}

	// Equal
	for i, a := range nums {
		for j, b := range nums {
			assertBool(t, i == j, Equal2(a, b))
		}
	}
}
