package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzCarry(t *testing.F) {

	r := func(env string, ctx int, a bool, b rune) string {
		switch {
		case env == "test":
			return ""
		case ctx == 1:
			return "1"
		case !a:
			return ""
		case ctx == 6:
			return "6"
		case b == '_':
			return ""
		case env == "prod":
			return env
		case b == '7':
			return "6"
		default:
			return "do it"
		}
	}

	// F = int -> bool -> rune -> Maybe[string]
	type F = Lambda[int, Lambda[bool, Lambda[rune, string]]]

	// R = string -> F
	type R = Lambda[string, F]

	var (
		f   R
		ff  Lambda[R, R]
		fff Lambda[Lambda[R, Lambda[string, F]], Lambda[R, Lambda[string, F]]]
	)

	// f :: R
	f = Carry4(r)

	// ff :: R -> R...
	// ff(f) = f
	ff = Carry2(Apply[string, F])

	// fff :: (R -> R...) -> R -> R...
	// fff(ff) = ff
	fff = FMap[R, Lambda[string, F], Lambda[R, Lambda[string, F]]]()

	t.Fuzz(func(t *testing.T, env string, ctx int, a bool, b rune) {
		want := r(env, ctx, a, b)

		assert.Equal(t, want, f(env)(ctx)(a)(b))

		a1 := Apply(f, env)
		assert.Equal(t, want, a1(ctx)(a)(b))

		a2 := Apply(a1, ctx)
		assert.Equal(t, want, a2(a)(b))

		a3 := Apply(a2, a)
		assert.Equal(t, want, a3(b))

		a4 := Apply(a3, b)
		assert.Equal(t, want, a4)

		assert.Equal(t, want, ff(f)(env)(ctx)(a)(b))
		assert.Equal(t, want, fff(ff)(f)(env)(ctx)(a)(b))
	})
}
