package church

type Boolean = Term

// True If Else == \If.\Else.If
func True(iF Term) Term {
	return func(eLse Term) Term { return iF }
}

// False If Else == \If.\Else.Else
func False(iF Term) Term {
	return Zero(iF)
}

/*
	Simplified Syntax:
		A ? B : C == \A.ABC
*/

// Not p == p ? False : True
func Not(p Boolean) Boolean {
	return p(False)(True)
}

// And p q == p ? q : p
func And(p Boolean) Term {
	return func(q Boolean) Boolean { return p(q)(p) }
}

// Or p q == p ? p : q
func Or(p Boolean) Term {
	return func(q Boolean) Boolean { return p(p)(q) }
}

// Xor p q == p ? not q : q
func Xor(p Boolean) Term {
	return func(q Term) Term { return p(Not(q))(q) }
}

// Pair x y p == \x.\y.\p.pxy
func Pair(left Term) Term {
	return func(right Term) Term {
		return func(p Term) Term {
			return p(left)(right)
		}
	}
}

// Left == \p . p True
// Left (Pair x y) = x
func Left(p Term) Term {
	return p(True)
}

// Right == \p . p False
// Right (Pair x y) = y
func Right(p Term) Term {
	return p(False)
}
