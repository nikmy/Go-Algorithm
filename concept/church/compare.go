package church

// IsZero n == n (\x.False) True
//
// explanation: if n equals Zero, it will take
// second argument (which is True), otherwise
// it will chain  (\x.False) ... (\x.False) True,
// which is always equal to False
func IsZero(n Numeral) Boolean {
	return n(func(Term) Term { return False })(True)
}

// LessOrEqual m n == IsZero (Sub m n)
func LessOrEqual(m Numeral) Term {
	return func(n Numeral) Boolean {
		return IsZero(Sub(m)(n))
	}
}

// Less m n == Not (LessOrEqual n m)
func Less(m Numeral) Term {
	return func(n Numeral) Boolean {
		return Not(LessOrEqual(n)(m))
	}
}

// Equal m n == And (LessOrEqual m n) (LessOrEqual n m)
func Equal(m Numeral) Term {
	return func(n Numeral) Boolean {
		return And(LessOrEqual(m)(n))(LessOrEqual(n)(m))
	}
}
