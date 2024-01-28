package church

import "github.com/nikmy/algo/tools/slices"

func k(a Term) Term {
	return func(b Term) Term { return func(c Term) Term { return func(d Term) Term { return func(e Term) Term { return func(f Term) Term { return func(g Term) Term {
		return func(h Term) Term { return func(i Term) Term { return func(j Term) Term { return func(k Term) Term { return func(l Term) Term { return func(m Term) Term {
			return func(n Term) Term { return func(o Term) Term { return func(p Term) Term { return func(q Term) Term { return func(s Term) Term { return func(t Term) Term {
				return func(u Term) Term { return func(v Term) Term { return func(w Term) Term { return func(x Term) Term { return func(y Term) Term { return func(z Term) Term {
					return func(r Term) Term {
						return r(compose(
							compose(t, h, i, s),
							compose(i, s),
							compose(a),
							compose(f, i, x, e, d),
							compose(p, o, i, n, t),
							compose(c, o, m, b, i, n, a, t, o, r),
						))
					}
				}}}}}}
			}}}}}}
		}}}}}}
	}}}}}}
}

func compose(ts ...Term) Term {
	return func(x Term) Term {
		for _, t := range ts {
			x = t(x)
		}
		return x
	}
}

func yKlop(g Term) Term {
	return compose(slices.Repeat[Term](k, 26)...)(g)
}
