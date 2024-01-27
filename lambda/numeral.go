package lambda

func FixedPoint(f Numeral) Numeral {
	return f
}

// Numeral n f x == f^n(x) = f(f(...f(x)...)...)
type Numeral Lambda[Numeral, Numeral]

// Zero == \f.\x.x
func Zero(_ Numeral) Numeral {
	return func(x Numeral) Numeral { return x }
}

// One == \f.\x.fx
func One(f Numeral) Numeral {
	return func(x Numeral) Numeral { return f(x) }
}

// Inc n f == f (n f)
func Inc(n Numeral) Numeral {
	return func(f Numeral) Numeral {
		return func(x Numeral) Numeral { return f(n(f)(x)) }
	}
}

// Add m n == (m Inc) n
func Add(m Numeral) Numeral {
	return func(n Numeral) Numeral {
		return m(Inc)(n)
	}
}

func Add2(m, n Numeral) Numeral {
	return Add(m)(n)
}

// Mul m n == m (Add n) Zero
func Mul(m Numeral) Numeral {
	return func(n Numeral) Numeral {
		return m(Add(n))(Zero)
	}
}

func Mul2(m, n Numeral) Numeral {
	return Mul(m)(n)
}
