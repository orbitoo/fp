package maybe

import "fmt"

type Maybe[T any] struct {
	value     T
	isNothing bool
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value, false}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{isNothing: true}
}

func (m Maybe[T]) Map(f func(T) T) Maybe[T] {
	if m.isNothing {
		return m
	}
	return Just(f(m.value))
}

func (m Maybe[T]) FlatMap(f func(T) Maybe[T]) Maybe[T] {
	if m.isNothing {
		return m
	}
	return f(m.value)
}

func (m Maybe[T]) Get() (T, bool) {
	return m.value, !m.isNothing
}

func (m Maybe[T]) IsNothing() bool {
	return m.isNothing
}

func (m Maybe[T]) String() string {
	if m.isNothing {
		return "Nothing"
	}
	return fmt.Sprintf("Just %v", m.value)
}

// XXX: an alternative way to implement applicative
func LiftMaybe1[T1 any, T any](f func(T1) T) func(Maybe[T1]) Maybe[T] {
	return func(m Maybe[T1]) Maybe[T] {
		if m.isNothing {
			return Nothing[T]()
		}
		return Just(f(m.value))
	}
}

func LiftMaybe2[T1 any, T2 any, T any](f func(T1, T2) T) func(Maybe[T1], Maybe[T2]) Maybe[T] {
	return func(m1 Maybe[T1], m2 Maybe[T2]) Maybe[T] {
		if m1.isNothing || m2.isNothing {
			return Nothing[T]()
		}
		return Just(f(m1.value, m2.value))
	}
}

func LiftMaybe3[T1 any, T2 any, T3 any, T any](f func(T1, T2, T3) T) func(Maybe[T1], Maybe[T2], Maybe[T3]) Maybe[T] {
	return func(m1 Maybe[T1], m2 Maybe[T2], m3 Maybe[T3]) Maybe[T] {
		if m1.isNothing || m2.isNothing || m3.isNothing {
			return Nothing[T]()
		}
		return Just(f(m1.value, m2.value, m3.value))
	}
}

func LiftMaybe4[T1 any, T2 any, T3 any, T4 any, T any](f func(T1, T2, T3, T4) T) func(Maybe[T1], Maybe[T2], Maybe[T3], Maybe[T4]) Maybe[T] {
	return func(m1 Maybe[T1], m2 Maybe[T2], m3 Maybe[T3], m4 Maybe[T4]) Maybe[T] {
		if m1.isNothing || m2.isNothing || m3.isNothing || m4.isNothing {
			return Nothing[T]()
		}
		return Just(f(m1.value, m2.value, m3.value, m4.value))
	}
}

func LiftMaybe5[T1 any, T2 any, T3 any, T4 any, T5 any, T any](f func(T1, T2, T3, T4, T5) T) func(Maybe[T1], Maybe[T2], Maybe[T3], Maybe[T4], Maybe[T5]) Maybe[T] {
	return func(m1 Maybe[T1], m2 Maybe[T2], m3 Maybe[T3], m4 Maybe[T4], m5 Maybe[T5]) Maybe[T] {
		if m1.isNothing || m2.isNothing || m3.isNothing || m4.isNothing || m5.isNothing {
			return Nothing[T]()
		}
		return Just(f(m1.value, m2.value, m3.value, m4.value, m5.value))
	}
}

// FIXME: applicative now is not implemented,
// due to the restriction of the type system in golang.
// func (m Maybe[T]) Apply(f Maybe[func(T) T]) Maybe[T] {
// 	if m.isNothing || f.isNothing {
// 		return Maybe[T]{isNothing: true}
// 	}
// 	return Just(f.value(m.value))
// }
