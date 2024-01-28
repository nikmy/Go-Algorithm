package church

// yCombPart == (\x.\y.y(xxy))
func yCombPart(x Term) Term {
	return func(y Term) Term {
		return y(x(x)(y))
	}
}

// yComb == (\x.\y.y(xxy))(\x.\y.y(xxy))
func yComb(g Term) Term {
	return (yCombPart)(yCombPart)(g)
}
