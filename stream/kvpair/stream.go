package kvpair

import (
	p "github.com/orbitoo/fp/pipe"
)

type pair[T1, T2 comparable] struct {
	key   T1
	value T2
}

type stream[T1, T2 comparable] struct {
	pipe *p.Pipe[*pair[T1, T2]]
}

func NewFromSlice[T1, T2 comparable](m map[T1]T2) *stream[T1, T2] {
	slice := make([]*pair[T1, T2], len(m))
	i := 0
	for k, v := range m {
		slice[i] = &pair[T1, T2]{k, v}
		i++
	}
	return &stream[T1, T2]{p.NewFromSlice(slice)}
}
