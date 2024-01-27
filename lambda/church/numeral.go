package church

// Numeral is such n that n f x == \f.f^n(x) = \f.f(f(...f(x)...)...)
type Numeral = Term

// Zero == \f.\x.x
func Zero(_ Term) Term {
	return func(x Term) Term { return x }
}

// One == \f.\x.fx
func One(f Term) Term {
	return func(x Term) Term { return f(x) }
}

// Inc n f == \f.\n.f(nf)
func Inc(n Numeral) Numeral {
	return func(f Term) Term {
		return func(x Term) Term { return f(n(f)(x)) }
	}
}

// Add m n == \m.\n.m Inc n
func Add(m Numeral) Term {
	return func(n Numeral) Numeral {
		return m(Inc)(n)
	}
}

// Mul m n == \m.\n.m (Add n) Zero
func Mul(m Numeral) Term {
	return func(n Numeral) Numeral {
		return m(Add(n))(Zero)
	}
}

/*
	Kleene's trick:

	Let's define incStep(N) as Nth iteration of { Pair _ n |-> Pair n (Inc n) },
	so it maps (Pair `0` `0`) to (Pair (`N-1` f) (`N` f)), and the left side
	of N(incStep)(Pair 0 0) is "decremented" N
*/
func incStep(p Term) Term {
	return Pair(Right(p))(Inc(Right(p)))
}

// Dec == \n.Left[ n incStep (Pair Zero Zero) ]
func Dec(n Numeral) Numeral {
	return Left(n(incStep)(Pair(Zero)(Zero)))
}

// Sub m n == m Dec n
func Sub(m Numeral) Term {
	return func(n Term) Numeral {
		return n(Dec)(m)
	}
}
