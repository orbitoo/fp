package function

func Curry2[T1 any, T2 any, T any](f func(T1, T2) T) func(T1) func(T2) T {
	return func(t1 T1) func(T2) T {
		return func(t2 T2) T {
			return f(t1, t2)
		}
	}
}

func Curry3[T1 any, T2 any, T3 any, T any](f func(T1, T2, T3) T) func(T1) func(T2) func(T3) T {
	return func(t1 T1) func(T2) func(T3) T {
		return func(t2 T2) func(T3) T {
			return func(t3 T3) T {
				return f(t1, t2, t3)
			}
		}
	}
}

func Curry4[T1 any, T2 any, T3 any, T4 any, T any](f func(T1, T2, T3, T4) T) func(T1) func(T2) func(T3) func(T4) T {
	return func(t1 T1) func(T2) func(T3) func(T4) T {
		return func(t2 T2) func(T3) func(T4) T {
			return func(t3 T3) func(T4) T {
				return func(t4 T4) T {
					return f(t1, t2, t3, t4)
				}
			}
		}
	}
}

func Curry5[T1 any, T2 any, T3 any, T4 any, T5 any, T any](f func(T1, T2, T3, T4, T5) T) func(T1) func(T2) func(T3) func(T4) func(T5) T {
	return func(t1 T1) func(T2) func(T3) func(T4) func(T5) T {
		return func(t2 T2) func(T3) func(T4) func(T5) T {
			return func(t3 T3) func(T4) func(T5) T {
				return func(t4 T4) func(T5) T {
					return func(t5 T5) T {
						return f(t1, t2, t3, t4, t5)
					}
				}
			}
		}
	}
}
