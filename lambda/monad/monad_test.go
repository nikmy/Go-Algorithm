package monad

func staticCheckMaybe[A any]() {
	_ = Return[A, Maybe[A]]()
	_ = Bind[A, Maybe[A]]()
}

func staticCheckResult[A any]() {
	_ = Return[A, Result[A]]()
	_ = Bind[A, Result[A]]()
}
