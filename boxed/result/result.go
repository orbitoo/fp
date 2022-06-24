package result

import "log"

type Result[T any] struct {
	v       T
	e       error
	isError bool
}

func Ok[T any](v T) Result[T] {
	return Result[T]{v: v}
}

func Err[T any](e error) Result[T] {
	return Result[T]{e: e, isError: true}
}

func (r Result[T]) Expect(s string) T {
	if r.isError {
		log.Println(s)
		log.Fatal(r.e)
	}
	return r.v
}

func (r Result[T]) Unwrap() T {
	if r.isError {
		log.Fatal(r.e)
	}
	return r.v
}

func (r Result[T]) Default(fn func() T) T {
	if r.isError {
		return fn()
	}
	return r.v
}

func (r Result[T]) Handle(fn func()) T {
	if r.isError {
		fn()
	}
	return r.v
}
