package church

import "testing"

func TestFact(t *testing.T) {
	f := 1

	// base
	assertInt(t, f, Fact(Zero))

	// step
	for i := 1; i < 6; i++ {
		f *= i
		assertInt(t, f, Fact(nums[i]))
	}
}

func TestFib(t *testing.T) {
	pp, p := 1, 1

	// base
	assertInt(t, pp, Fib(Zero))
	assertInt(t, p, Fib(One))

	// step
	for _, n := range nums[2:20] {
		pp, p = p, pp+p
		assertInt(t, p, Fib(n))
	}
}

func TestIsPrime(t *testing.T) {
	primes := make([]bool, len(nums))
	for i := 2; i < len(nums); i++ {
		primes[i] = true
		for d := 2; d*d <= i; d++ {
			if i%d == 0 {
				primes[i] = false
				break
			}
		}
	}

	for i, n := range nums {
		assertBool(t, primes[i], IsPrime(n), "%d check fail", i)
	}
}
