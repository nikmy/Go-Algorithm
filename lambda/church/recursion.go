package church

// Y g == (\x.\y.g(xx)y)) (\x.\y.g(xx)y)
//
/*
	Y-combinator — recursion implementation in lambda calculus.

	For any combinator G:

		Y G = ( \g.(\x.\y.g(xx)y))(\x.\y.g(xx)y) ) G
			= (\x.\y.G(xx)y))(\x.\y.G(xx)y)
			= G((\x.\y.G(xx)y)(\x.\y.G(xx)y))
			= G(Y G)

	Usage: to avoid stack overflow, wrap recursive call R(f):

		Y(
			func(f Term) Term {
				return func(n Numeral) Term {
					return Predicate(n)(
						func(_ Term) Term { return One  })(
						func(_ Term) Term { return R(f) },
					)(nil)
				}
			},
		)

	This way, recursion is not infinite because of deferred execution
	of recursive part. So, go will not evaluate it before direct call.

*/
func Y(g Term) Term {
	return (func(x Term) Term {
		return func(y Term) Term { return g(x(x))(y) }
	})(func(x Term) Term {
		return func(y Term) Term { return g(x(x))(y) }
	})
}


func genFact(fact Term) Term {
	return func(n Numeral) Numeral {
		return IsZero(Dec(n))(
			func(_ Term) Numeral { return One })(
			func(_ Term) Numeral { return Mul(n)(fact(Dec(n))) },
		)(nil)
	}
}

func Fact(n Numeral) Numeral {
	return Y(genFact)(n)
}

func genFib(fib Term) Term {
	return func(n Numeral) Numeral {
		return LessOrEqual(n)(One)(
			func(_ Term) Numeral { return One })(
			func(_ Term) Numeral { return Add(fib(Dec(n)))(fib(Dec(Dec(n)))) },
		)(nil)
	}
}

func Fib(n Numeral) Numeral {
	return Y(genFib)(n)
}
