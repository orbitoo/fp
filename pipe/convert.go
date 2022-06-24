package pipe

type PipeConvert[T1, T2 comparable] struct {
	in   chan T1
	out  chan T2
	done chan bool
}

func (p *PipeConvert[T1, T2]) Calc(fn func(x T1) T2) {
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

func (p *PipeConvert[T1, T2]) Connect() *Pipe[T2] {
	return &Pipe[T2]{
		in:   p.out,
		out:  make(chan T2),
		done: p.done,
	}
}

func (p *PipeConvert[T1, T2]) Out() chan T2 {
	return p.out
}
