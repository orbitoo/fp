package kvpair

func (s *stream[T1, T2]) Collect() map[T1]T2 {
	r := make(map[T1]T2)
	for v := range s.pipe.Out() {
		r[v.key] = v.value
	}
	return r
}

func (s *stream[T1, T2]) FirstK(cond func(x T1) bool) (T1, T2, bool) {
	s.pipe = s.pipe.Connect()
	s = s.FilterK(cond)
	return s.Get()
}

func (s *stream[T1, T2]) Get() (T1, T2, bool) {
	r, ok := <-s.pipe.Out()
	return r.key, r.value, ok
}
