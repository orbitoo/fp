package slice

func (s *stream[T]) Filter(cond func(x T) bool) *stream[T] {
	s.pipe = s.pipe.Connect()
	s.pipe.Filter(cond)
	return s
}

func (s *stream[T]) Drop(cond func(x T) bool) *stream[T] {
	s.pipe = s.pipe.Connect()
	fn := func(x T) bool { return !cond(x) }
	s.pipe.Filter(fn)
	return s
}

func (s *stream[T]) Map(mapping func(x T) T) *stream[T] {
	s.pipe = s.pipe.Connect()
	s.pipe.Calc(mapping)
	return s
}

func (s *stream[T]) IfElse(cond func(x T) bool, fn1 func(x T) T, fn2 func(x T) T) *stream[T] {
	s.pipe = s.pipe.Connect()
	fn := func(x T) T {
		if cond(x) {
			return fn1(x)
		} else {
			return fn2(x)
		}
	}
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T]) If(cond func(x T) bool, fn func(x T) T) *stream[T] {
	s.pipe = s.pipe.Connect()
	f := func(x T) T {
		if cond(x) {
			return fn(x)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T]) Unless(cond func(x T) bool, fn func(x T) T) *stream[T] {
	s.pipe = s.pipe.Connect()
	f := func(x T) T {
		if !cond(x) {
			return fn(x)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}
