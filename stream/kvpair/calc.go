package kvpair

func (s *stream[T1, T2]) FilterK(cond func(x T1) bool) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) bool { return cond(x.key) }
	s.pipe.Filter(fn)
	return s
}

func (s *stream[T1, T2]) FilterV(cond func(x T2) bool) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) bool { return cond(x.value) }
	s.pipe.Filter(fn)
	return s
}

func (s *stream[T1, T2]) DropK(cond func(x T1) bool) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) bool { return !cond(x.key) }
	s.pipe.Filter(fn)
	return s
}

func (s *stream[T1, T2]) DropV(cond func(x T2) bool) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) bool { return !cond(x.value) }
	s.pipe.Filter(fn)
	return s
}

func (s *stream[T1, T2]) MapK(mapping func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] { x.key = mapping(x.key); return x }
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) MapV(mapping func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] { x.value = mapping(x.value); return x }
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) IfElseKK(cond func(x T1) bool, fn1 func(x T1) T1, fn2 func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.key) {
			x.key = fn1(x.key)
		} else {
			x.key = fn2(x.key)
		}
		return x
	}
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) IfElseKV(cond func(x T1) bool, fn1 func(x T2) T2, fn2 func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.key) {
			x.value = fn1(x.value)
		} else {
			x.value = fn2(x.value)
		}
		return x
	}
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) IfElseVK(cond func(x T2) bool, fn1 func(x T1) T1, fn2 func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.value) {
			x.key = fn1(x.key)
		} else {
			x.key = fn2(x.key)
		}
		return x
	}
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) IfElseVV(cond func(x T2) bool, fn1 func(x T2) T2, fn2 func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	fn := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.value) {
			x.value = fn1(x.value)
		} else {
			x.value = fn2(x.value)
		}
		return x
	}
	s.pipe.Calc(fn)
	return s
}

func (s *stream[T1, T2]) IfKK(cond func(x T1) bool, fn func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.key) {
			x.key = fn(x.key)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) IfKV(cond func(x T1) bool, fn func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.key) {
			x.value = fn(x.value)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) IfVK(cond func(x T2) bool, fn func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.value) {
			x.key = fn(x.key)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) IfVV(cond func(x T2) bool, fn func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if cond(x.value) {
			x.value = fn(x.value)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) UnlessKK(cond func(x T1) bool, fn func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if !cond(x.key) {
			x.key = fn(x.key)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) UnlessKV(cond func(x T1) bool, fn func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if !cond(x.key) {
			x.value = fn(x.value)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) UnlessVK(cond func(x T2) bool, fn func(x T1) T1) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if !cond(x.value) {
			x.key = fn(x.key)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}

func (s *stream[T1, T2]) UnlessVV(cond func(x T2) bool, fn func(x T2) T2) *stream[T1, T2] {
	s.pipe = s.pipe.Connect()
	f := func(x *pair[T1, T2]) *pair[T1, T2] {
		if !cond(x.value) {
			x.value = fn(x.value)
		}
		return x
	}
	s.pipe.Calc(f)
	return s
}
