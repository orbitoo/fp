package slice

import (
	p "github.com/orbitoo/fp/pipe"
)

type stream[T comparable] struct {
	pipe *p.Pipe[T]
}

func NewFromSlice[T comparable](slice []T) *stream[T] {
	return &stream[T]{p.NewFromSlice(slice)}
}

func NewFromIteration[T comparable](fn func(x T) T, x T) *stream[T] {
	return &stream[T]{p.NewFromIteration(fn, x)}
}
