package pipe

type Pipe[T comparable] struct {
	in   chan T
	out  chan T
	done chan bool
}

func NewFromSlice[T comparable](slice []T) *Pipe[T] {
	p := &Pipe[T]{
		out:  make(chan T),
		done: make(chan bool),
	}
	done := p.done
	go func() {
		for _, v := range slice {
			select {
			case p.out <- v:
			case <-done:
				close(p.out)
				return
			}
		}
		close(p.out)
	}()
	return p
}

func NewFromIteration[T comparable](fn func(x T) T, x T) *Pipe[T] {
	p := &Pipe[T]{
		out:  make(chan T),
		done: make(chan bool),
	}
	done := p.done
	go func() {
		for {
			select {
			case p.out <- x:
				x = fn(x)
			case <-done:
				close(p.out)
				return
			}
		}
	}()
	return p
}

func (p *Pipe[T]) Calc(fn func(x T) T) {
	go func() {
		for v := range p.in {
			select {
			case p.out <- fn(v):
			case <-p.done:
				close(p.out)
				return
			}
		}
		close(p.out)
	}()
}

func (p *Pipe[T]) Filter(cond func(x T) bool) {
	go func() {
		for v := range p.in {
			if cond(v) {
				select {
				case p.out <- v:
				case <-p.done:
					close(p.out)
					return
				}
			}
		}
		close(p.out)
	}()
}

func (p *Pipe[T]) Pick(n int) {
	done := p.done
	p.done = make(chan bool)
	go func() {
		i := 0
		for v := range p.in {
			p.out <- v
			i++
			if i == n {
				close(done)
				close(p.out)
				return
			}
		}
	}()
}

func (p *Pipe[T]) PickWhile(cond func(x T) bool) {
	done := p.done
	p.done = make(chan bool)
	go func() {
		i := 0
		for v := range p.in {
			if cond(v) {
				p.out <- v
				i++
			} else {
				close(done)
				close(p.out)
				return
			}
		}
	}()
}

func (p *Pipe[T]) Skip(n int) {
	go func() {
		i := 0
		for range p.in {
			i++
			if i == n {
				return
			}
		}
	}()
}

func (p *Pipe[T]) SkipWhile(cond func(x T) bool) {
	go func() {
		i := 0
		for v := range p.in {
			if !cond(v) {
				p.out <- v
				i++
			}
		}
		close(p.out)
	}()
}

func (p *Pipe[T]) Connect() *Pipe[T] {
	return &Pipe[T]{
		in:   p.out,
		out:  make(chan T),
		done: p.done,
	}
}

func (p *Pipe[T]) Out() chan T {
	return p.out
}

func (p *Pipe[T]) End() {
	close(p.done)
}
