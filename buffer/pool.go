package buffer

import (
	"go.uber.org/zap/internal/pool"
)

type Pool struct {
	p *pool.Pool[*Buffer]
}

func NewPool() Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (p Pool) Get() *Buffer { _ = "STUB: not implemented"; return nil }

func (p Pool) put(buf *Buffer) { _ = "STUB: not implemented"; return }
