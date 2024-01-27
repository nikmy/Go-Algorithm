package lambda

func Const[T any](value T) func() T {
	return func() T { return value }
}

type Unit = struct{}

func Do(action Lambda[Unit, Unit]) {
	_ = action(Unit{})
}

func NewAction[T any](f func()) Lambda[Unit, Unit] {
	return func(Unit) Unit { f(); return Unit{} }
}

func NewNoArg[R any](f func() R) Lambda[Unit, R] {
	return func(Unit) R { return f() }
}

func NewNoReturn[Arg any](f func(Arg)) Lambda[Arg, Unit] {
	return func(x Arg) Unit { f(x); return Unit{} }
}

type Lambda[Arg, Result any] func(Arg) Result

func Apply[Arg, Result any](lambda Lambda[Arg, Result], arg Arg) Result {
	return lambda(arg)
}

func FMap[Arg, Result any, F Lambda[Arg, Result]]() Lambda[F, Lambda[Arg, Result]] {
	return func(lambda F) Lambda[Arg, Result] {
		return func(arg Arg) Result {
			return lambda(arg)
		}
	}
}
