package slice

import "log"

func (s *stream[T]) Pick(n int) *stream[T] {
	if n <= 0 {
		log.Printf("n must be greater than 0")
		log.Printf("If you want to end the stream, use End()")
		return s
	}
	s.pipe = s.pipe.Connect()
	s.pipe.Pick(n)
	return s
}

func (s *stream[T]) PickWhile(cond func(x T) bool) *stream[T] {
	s.pipe = s.pipe.Connect()
	s.pipe.PickWhile(cond)
	return s
}

func (s *stream[T]) Skip(n int) *stream[T] {
	s.pipe = s.pipe.Connect()
	s.pipe.Skip(n)
	return s
}

func (s *stream[T]) SkipWhile(cond func(x T) bool) *stream[T] {
	s.pipe = s.pipe.Connect()
	s.pipe.SkipWhile(cond)
	return s
}
