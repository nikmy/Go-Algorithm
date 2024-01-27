package lambda

func Carry2[Arg1, Arg2, Result any](f func(Arg1, Arg2) Result) Lambda[Arg1, Lambda[Arg2, Result]] {
	return func(x Arg1) Lambda[Arg2, Result] {
		return func(y Arg2) Result {
			return f(x, y)
		}
	}
}

func Carry3[
	Arg1, Arg2, Arg3,
	Result any,
](
	f func(Arg1, Arg2, Arg3) Result,
) Lambda[Arg1, Lambda[Arg2, Lambda[Arg3, Result]]] {

	return func(x Arg1) Lambda[Arg2, Lambda[Arg3, Result]] {
		return func(y Arg2) Lambda[Arg3, Result] {
			return func(z Arg3) Result {
				return f(x, y, z)
			}
		}
	}
}

func Carry4[
	Arg1, Arg2, Arg3, Arg4,
	Result any,
](
	f func(Arg1, Arg2, Arg3, Arg4) Result,
) Lambda[Arg1, Lambda[Arg2, Lambda[Arg3, Lambda[Arg4, Result]]]] {

	return func(x Arg1) Lambda[Arg2, Lambda[Arg3, Lambda[Arg4, Result]]] {
		return func(y Arg2) Lambda[Arg3, Lambda[Arg4, Result]] {
			return func(z Arg3) Lambda[Arg4, Result] {
				return func(t Arg4) Result {
					return f(x, y, z, t)
				}
			}
		}
	}
}
