package kvpair

func (s *stream[T1, T2]) Pick(n int) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	s.pipe.Pick(n)
	return s
}

func (s *stream[T1, T2]) Skip(n int) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	s.pipe.Skip(n)
	return s
}
