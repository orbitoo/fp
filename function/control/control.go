package control

type IfElse[T any] struct {
	result T
	done   bool
}

func Ternary[T any](cond bool, v1, v2 T) T {
	if cond {
		return v1
	} else {
		return v2
	}
}

func If[T any](cond bool, v T) *IfElse[T] {
	if cond {
		return &IfElse[T]{v, true}
	} else {
		return &IfElse[T]{v, false}
	}
}

func (ie *IfElse[T]) ElseIf(cond bool, v T) *IfElse[T] {
	if !ie.done && cond {
		ie.result = v
		ie.done = true
	}
	return ie
}

func (ie *IfElse[T]) Else(v T) T {
	if ie.done {
		return ie.result
	} else {
		return v
	}
}

func For(n int, f func(index int)) {
	for i := 0; i < n; i++ {
		f(i)
	}
}

func ForReverse(n int, f func(index int)) {
	for i := n - 1; i >= 0; i-- {
		f(i)
	}
}

func While(cond func() bool, f func()) {
	for cond() {
		f()
	}
}
