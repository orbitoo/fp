package boxed

import "fmt"

type Either[T1, T2 any] struct {
	left   T1
	right  T2
	isLeft bool
}

func Left[T1, T2 any](value T1) Either[T1, T2] {
	return Either[T1, T2]{left: value, isLeft: true}
}

func Right[T1, T2 any](value T2) Either[T1, T2] {
	return Either[T1, T2]{right: value}
}

func (e Either[T1, T2]) Map(f func(T2) T2) Either[T1, T2] {
	if e.isLeft {
		return Either[T1, T2]{left: e.left, isLeft: true}
	}
	return Either[T1, T2]{right: f(e.right)}
}

func (e Either[T1, T2]) FlatMap(f func(T2) Either[T1, T2]) Either[T1, T2] {
	if e.isLeft {
		return Either[T1, T2]{left: e.left, isLeft: true}
	}
	return f(e.right)
}

func (e Either[T1, T2]) Get() (T1, T2, bool) {
	return e.left, e.right, !e.isLeft
}

func (e Either[T1, T2]) IsLeft() bool {
	return e.isLeft
}

func (e Either[T1, T2]) String() string {
	if e.isLeft {
		return fmt.Sprintf("Left %v", e.left)
	}
	return fmt.Sprintf("Right %v", e.right)
}

// In haskell, applicative is instance of Either L.
func LiftEither1[L, R, T1 any](f func(T1) R) func(Either[L, T1]) Either[L, R] {
	return func(e Either[L, T1]) Either[L, R] {
		if e.isLeft {
			return Either[L, R]{left: e.left, isLeft: true}
		}
		return Right[L](f(e.right))
	}
}

func LiftEither2[L, R, T1 any, T2 any](f func(T1, T2) R) func(Either[L, T1], Either[L, T2]) Either[L, R] {
	return func(e1 Either[L, T1], e2 Either[L, T2]) Either[L, R] {
		if e1.isLeft {
			return Either[L, R]{left: e1.left, isLeft: true}
		} else if e2.isLeft {
			return Either[L, R]{left: e2.left, isLeft: true}
		}
		return Right[L](f(e1.right, e2.right))
	}
}

func LiftEither3[L, R, T1 any, T2 any, T3 any](f func(T1, T2, T3) R) func(Either[L, T1], Either[L, T2], Either[L, T3]) Either[L, R] {
	return func(e1 Either[L, T1], e2 Either[L, T2], e3 Either[L, T3]) Either[L, R] {
		if e1.isLeft {
			return Either[L, R]{left: e1.left, isLeft: true}
		} else if e2.isLeft {
			return Either[L, R]{left: e2.left, isLeft: true}
		} else if e3.isLeft {
			return Either[L, R]{left: e3.left, isLeft: true}
		}
		return Right[L](f(e1.right, e2.right, e3.right))
	}
}

func LiftEither4[L, R, T1 any, T2 any, T3 any, T4 any](f func(T1, T2, T3, T4) R) func(Either[L, T1], Either[L, T2], Either[L, T3], Either[L, T4]) Either[L, R] {
	return func(e1 Either[L, T1], e2 Either[L, T2], e3 Either[L, T3], e4 Either[L, T4]) Either[L, R] {
		if e1.isLeft {
			return Either[L, R]{left: e1.left, isLeft: true}
		} else if e2.isLeft {
			return Either[L, R]{left: e2.left, isLeft: true}
		} else if e3.isLeft {
			return Either[L, R]{left: e3.left, isLeft: true}
		} else if e4.isLeft {
			return Either[L, R]{left: e4.left, isLeft: true}
		}
		return Right[L](f(e1.right, e2.right, e3.right, e4.right))
	}
}

func LiftEither5[L, R, T1 any, T2 any, T3 any, T4 any, T5 any](f func(T1, T2, T3, T4, T5) R) func(Either[L, T1], Either[L, T2], Either[L, T3], Either[L, T4], Either[L, T5]) Either[L, R] {
	return func(e1 Either[L, T1], e2 Either[L, T2], e3 Either[L, T3], e4 Either[L, T4], e5 Either[L, T5]) Either[L, R] {
		if e1.isLeft {
			return Either[L, R]{left: e1.left, isLeft: true}
		} else if e2.isLeft {
			return Either[L, R]{left: e2.left, isLeft: true}
		} else if e3.isLeft {
			return Either[L, R]{left: e3.left, isLeft: true}
		} else if e4.isLeft {
			return Either[L, R]{left: e4.left, isLeft: true}
		} else if e5.isLeft {
			return Either[L, R]{left: e5.left, isLeft: true}
		}
		return Right[L](f(e1.right, e2.right, e3.right, e4.right, e5.right))
	}
}
