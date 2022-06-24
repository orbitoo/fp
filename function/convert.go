package function

func Dot[T, T1, T2 any](f func(T) T2, g func(T1) T) func(T1) T2 {
	return func(x T1) T2 {
		return f(g(x))
	}
}

func Flip[T1, T2, T3 any](f func(x T1, y T2) T3) func(T2, T1) T3 {
	return func(y T2, x T1) T3 {
		return f(x, y)
	}
}
