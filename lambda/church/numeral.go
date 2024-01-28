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

/*
	Tricky one is division. Recursive definition is

		div(m, n) == (m <= n) ? 0 : 1 + div(m - n, n),

	or in lambda notation

		Div m n == (Less m, n) Zero Inc(Div (Sub m n) n) (1).

	We cannot use built-in recursion, so we need to find
	non-recursive representation of Div.

	If we let Div be a variable in (1), we can rewrite (1) as

		Div = genDiv Div,

	where

		genDiv G = LessOrEqual(m, n) Zero Inc(G (Sub m n) n).

	So, if a solution exists, Div is a fixed point of genDiv
	combinator. To find a fixed point, we can use powerful
	tool Y-combinator:

		YComb == (\x.\y.y(xxy))(\x.\y.y(xxy))

	For any combinator F:

		YComb F = (\x.\y.y(xxy)) (\x.\y.y(xxy)) F
				= F((\x.\y.y(xxy)) (\x.\y.y(xxy)) F)
				= F(YComb F)

	(if you don't understand the first transition, just
	substitute (\x.\y.y(xxy)) as x and F as y into the first
	closure).

	So, now we have a formula for Div:

		Div == YComb (\G . Less(m, n) Zero Inc(G (Sub m n) n))
*/

// yCombPart == (\x.\y.y(xxy))
func yCombPart(x Term) Term {
	return func(y Term) Term {
		return y(x(x)(y))
	}
}

// YComb == (\x.\y.y(xxy))(\x.\y.y(xxy))
func YComb(g Term) Term {
	return (yCombPart)(yCombPart)(g)
}

// genDiv G == Less(m, n) Zero Inc(G (Sub m n) n)
func genDiv(g Term) Term {
	return func(m Numeral) Term {
		return func(n Numeral) Term {
			return (Less(m)(n))(Zero)(Inc(g(Sub(m)(n))(n)))
		}
	}
}

// div == YComb genDiv
func div(m Numeral) Term {
	return YComb(genDiv)(m)
}

var _ = div // unfortunately, it does not actually work :(

// Div dirty implementation
//
// We have to use dirty hack to bring laziness
// to our computations, otherwise it will
// overflow callstack (because go runtime
// evaluates all arguments before apply).
//
// Test for this function does not check YComb
// trick, but it checks recursive formula.
func Div(m Numeral) Term {
	return func(n Term) Numeral {
		return ifThenElse(Less(m)(n), Zero, Inc(Div(Sub(m)(n))(n)))
	}
}

// Mod dirty implementation
//
// Mod m n == (Less m n) m Mod(Sub m n)
func Mod(m Numeral) Term {
	return func(n Numeral) Term {
		return ifThenElse(Less(m)(n), m, Mod(Sub(m)(n))(n))
	}
}
