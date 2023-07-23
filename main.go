package workers

import (
	"errors"

	"sync/atomic"
)

type Pool struct {
	size    int32
	counter *atomic.Int32
}

func NewPool(poolSize int32) *Pool {
	p := &Pool{
		size:    poolSize,
		counter: &atomic.Int32{},
	}
	return p
}

func (p *Pool) Woker(f func(...any)) error {
	if p.counter.Load() == p.size {
		return errors.New("all workers are busy right now")
	}
	p.counter.Add(1)
	go func() {
		f()
		p.counter.Add(-1)
	}()
	return nil
}
