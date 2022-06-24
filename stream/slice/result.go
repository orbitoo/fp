package slice

func (s *stream[T]) Collect() []T {
	r := make([]T, 0)
	for v := range s.pipe.Out() {
		r = append(r, v)
	}
	return r
}

func (s *stream[T]) First(cond func(x T) bool) (T, bool) {
	s = s.Filter(cond)
	return s.Get()
}

func (s *stream[T]) Get() (T, bool) {
	r, ok := <-s.pipe.Out()
	return r, ok
}

func (s *stream[T]) ForEach(fn func(x T)) {
	for v := range s.pipe.Out() {
		fn(v)
	}
}

func (s *stream[T]) Reduce(fn func(x, y T) T, x T) T {
	res := x
	for v := range s.pipe.Out() {
		res = fn(res, v)
	}
	return res
}

func (s *stream[T]) End() {
	s.pipe.End()
}
