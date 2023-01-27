package packetio

import "sync"

type Cond struct {
	mu        *sync.Mutex
	cond      *sync.Cond
	condCount int
}

func NewCond(mu *sync.Mutex) *Cond {
	return &Cond{
		mu:   mu,
		cond: sync.NewCond(mu),
	}
}

func (c *Cond) Wait() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for {
		if c.condCount > 0 {
			c.condCount--
			return
		}
		c.cond.Wait()
	}
}

func (c *Cond) Signal() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.condCount++
	c.cond.Signal()
}
