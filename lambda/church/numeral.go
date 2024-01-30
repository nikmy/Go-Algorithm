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
   so it maps (Pair `0` `0`) to (Pair `N-1` `N`), and the left side
   of N(incStep)(Pair `0` `0`) is "decremented" N
*/

// incStep p == Pair (Right p) (Inc (Right p))
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

/*
   Tricky one is division. Recursive definition is

       div(m, n) == (m < n) ? 0 : 1 + div(m - n, n),

   or in lambda notation

       Div m n == (Less m n) Zero Inc(Div (Sub m n) n) (1).

   We cannot use built-in recursion, so we need to find
   non-recursive representation of Div.

   If we let Div be a variable in (1), we can rewrite (1) as

       Div m n = (genDivBy n Div) m,

   where

       genDivBy n G = \m . (Less m n) Zero Inc(G (Sub m n)).

   Note: we need function of one argument to be passed to Y-combinator.
   Since we need n in place inside and outside G's call, we need to
   rearrange arguments.

   So, if a solution exists, Div is a fixed point of genDiv
   combinator. To find a fixed point, we can use powerful
   tool Y-combinator (see recursion.go):

       Y == \g.(\x.\y.g(xx)y))(\x.\y.g(xx)y)

   So, now we have a formula for Div:

       Div m n == Y (genDivBy n) m
*/

// genDivBy n == \G. \m . LessOrEqual(m, n) Zero Inc(G (Sub m n))
func genDivBy(n Numeral) Term {
	return func(d Term) Term {
		return func(m Numeral) Numeral {
			return Less(m)(n)(
				func(_ Term) Numeral { return Zero })(
				func(_ Term) Numeral { return Inc(d(Sub(m)(n))) },
			)(nil)
		}
	}
}

// Div m n == Y (genDivBy n) m
func Div(m Numeral) Term {
	return func(n Numeral) Numeral {
		return Y(genDivBy(n))(m)
	}
}

// genModBy n == \G . \m . (Less m n) m (G (Sub m n))
func genModBy(n Numeral) Term {
	return func(g Term) Term {
		return func(m Numeral) Term {
			return Less(m)(n)(
				func(_ Term) Term { return m })(
				func(_ Term) Term { return g(Sub(m)(n)) },
			)(nil)
		}
	}
}

// Mod m n == Y (genModBy n) m
func Mod(m Numeral) Term {
	return func(n Numeral) Numeral {
		return Y(genModBy(n))(m)
	}
}
