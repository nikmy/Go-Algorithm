package monad

import . "github.com/nikmy/algo/lambda/function"

func Return[
	T any,
	M interface{
		make(T) M[T]
	},
]() Lambda[T, M[T]] {
	return func(x T) M[T] {
		var m M
		return m.make(x)
	}
}

func Bind[
	A any,
	M interface{
		make(A) M[A]
		chain(Lambda[A, M[A]]) M[A]
	},
]() Lambda[M[A], Lambda[Lambda[A, M[A]], M[A]]] {
	return func(ma M) Lambda[Lambda[A, M], M] {
		return func(f Lambda[A, M]) M { return ma.chain(f) }
	}
}
