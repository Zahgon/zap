package zapcore

import (
	"bufio"
	"sync"
	"time"
)

const (
	_defaultBufferSize = 256 * 1024

	_defaultFlushInterval = 30 * time.Second
)

type BufferedWriteSyncer struct {
	WS WriteSyncer

	Size int

	FlushInterval time.Duration

	Clock Clock

	mu          sync.Mutex
	initialized bool
	stopped     bool
	writer      *bufio.Writer
	ticker      *time.Ticker
	stop        chan struct{}
	done        chan struct{}
}

func (s *BufferedWriteSyncer) initialize() { _ = "STUB: not implemented"; return }

func (s *BufferedWriteSyncer) Write(bs []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *BufferedWriteSyncer) Sync() error { _ = "STUB: not implemented"; return nil }

func (s *BufferedWriteSyncer) flushLoop() { _ = "STUB: not implemented"; return }

func (s *BufferedWriteSyncer) Stop() (err error) { _ = "STUB: not implemented"; return nil }
