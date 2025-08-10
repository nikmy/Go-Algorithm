package church

/*
   Y-combinator — recursion implementation in lambda calculus.

   For any combinator G:

       Y G = (\x.\y.G(xx)y))(\x.\y.G(xx)y)
           = G((\x.\y.G(xx)y)(\x.\y.G(xx)y))
           = G(Y G)

   Usage: to avoid stack overflow, wrap recursive call R(f):

       Y(func(f Term) Term {
           return func(n Numeral) Term {
               return Predicate(n)(
                   func(_ Term) Term { return One })(
                   func(_ Term) Term { return R(f) },
               )(nil)
           }
       })

   This way recursion is not infinite because of deferred execution
   of recursive part. So, Go will not evaluate it before direct call.
*/

// Y g == (\x.\y.g(xx)y)) (\x.\y.g(xx)y)
func Y(g Term) Term {
	return (func(x Term) Term {
		return func(y Term) Term { return g(x(x))(y) }
	})(func(x Term) Term {
		return func(y Term) Term { return g(x(x))(y) }
	})
}

// genFact f == \n . IsZero(Dec n) One (Mul n (f (Dec n)))
func genFact(fact Term) Term {
	return func(n Numeral) Numeral {
		return IsZero(Dec(n))(
			func(_ Term) Numeral { return One })(
			func(_ Term) Numeral { return Mul(n)(fact(Dec(n))) },
		)(nil)
	}
}

// Fact == Y genFact
func Fact(n Numeral) Numeral {
	return Y(genFact)(n)
}

// genFib f == \n . (LessOrEqual n One) One (Add (f (Dec n)) (f (Dec (Dec n))))
func genFib(fib Term) Term {
	return func(n Numeral) Numeral {
		return LessOrEqual(n)(One)(
			func(_ Term) Numeral { return One })(
			func(_ Term) Numeral { return Add(fib(Dec(n)))(fib(Dec(Dec(n)))) },
		)(nil)
	}
}

// Fib == Y genFib
func Fib(n Numeral) Numeral {
	return Y(genFib)(n)
}

/*
   Prime check with O(sqrt(N)) mod operations

   We want to check each x from 2...ceil(sqrt n)
   whether it is divisor of n or not. "Cycle"
   implementation is recursive (pseudocode):

   isNotDiv(x):
       if (x * x > n)
           then return True
           else if (n % x == 0)
               then return False
               else return isNotDiv(Inc x)

   Note: we will start with x = 2 and end on x = ceil(sqrt(n)).
   Now we can use this check to check whether n is prime.
*/

// genIsPrimeStep n check == \d . (Less n (Mul d d)) True (IsZero(Mod n d) False check(Inc d))
func genIsPrimeStep(n Numeral) Term {
	return func(isPrimeStep Term) Term {
		return func(divisor Term) Term {
			return Less(n)(Mul(divisor)(divisor))(
				func(_ Term) Boolean { return True })(
				func(_ Term) Boolean { return IsZero(Mod(n)(divisor))(False)(isPrimeStep(Inc(divisor))) },
			)(nil)
		}
	}
}

// checkDivisors n == Y (genIsPrimeStep n) (Inc One)
func checkDivisors(n Numeral) Boolean {
	return Y(genIsPrimeStep(n))(Inc(One))
}

// IsPrime n == (IsZero (Dec n)) False (checkDivisors n)
func IsPrime(n Numeral) Boolean {
	return IsZero(Dec(n))(False)(checkDivisors(n))
}
