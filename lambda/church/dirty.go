package church

func makeInt(i Numeral) int {
	var (
		depth   = 0
		counter = func(Numeral) Numeral { depth++; return nil }
	)
	i(counter)(nil)
	return depth
}

func makeBool(b Boolean) bool {
	var (
		value bool
		mark  Term = func(Term) Term { value = true; return nil }
	)
	b(mark)(Zero)(nil)
	return value
}

func ifThenElse(predicate Boolean, iF Term, eLse Term) Term {
	if makeBool(predicate) {
		return iF
	}
	return eLse
}
