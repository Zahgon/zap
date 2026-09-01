package zapcore

import (
	"io"
	"sync"
)

type WriteSyncer interface {
	io.Writer
	Sync() error
}

func AddSync(w io.Writer) WriteSyncer { _ = "STUB: not implemented"; return *new(WriteSyncer) }

type lockedWriteSyncer struct {
	sync.Mutex
	ws WriteSyncer
}

func Lock(ws WriteSyncer) WriteSyncer { _ = "STUB: not implemented"; return *new(WriteSyncer) }

func (s *lockedWriteSyncer) Write(bs []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *lockedWriteSyncer) Sync() error { _ = "STUB: not implemented"; return nil }

type writerWrapper struct {
	io.Writer
}

func (w writerWrapper) Sync() error { _ = "STUB: not implemented"; return nil }

type multiWriteSyncer []WriteSyncer

func NewMultiWriteSyncer(ws ...WriteSyncer) WriteSyncer {
	_ = "STUB: not implemented"
	return *new(WriteSyncer)
}

func (ws multiWriteSyncer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (ws multiWriteSyncer) Sync() error { _ = "STUB: not implemented"; return nil }
