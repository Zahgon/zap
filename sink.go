package zap

import (
	"io"
	"net/url"
	"os"
	"sync"

	"go.uber.org/zap/zapcore"
)

const schemeFile = "file"

var _sinkRegistry = newSinkRegistry()

type Sink interface {
	zapcore.WriteSyncer
	io.Closer
}

type errSinkNotFound struct {
	scheme string
}

func (e *errSinkNotFound) Error() string { _ = "STUB: not implemented"; return "" }

type nopCloserSink struct{ zapcore.WriteSyncer }

func (nopCloserSink) Close() error { _ = "STUB: not implemented"; return nil }

type sinkRegistry struct {
	mu        sync.Mutex
	factories map[string]func(*url.URL) (Sink, error)
	openFile  func(string, int, os.FileMode) (*os.File, error)
}

func newSinkRegistry() *sinkRegistry { _ = "STUB: not implemented"; return nil }

func (sr *sinkRegistry) RegisterSink(scheme string, factory func(*url.URL) (Sink, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *sinkRegistry) newSink(rawURL string) (Sink, error) {
	_ = "STUB: not implemented"
	return *new(Sink), nil
}

func RegisterSink(scheme string, factory func(*url.URL) (Sink, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *sinkRegistry) newFileSinkFromURL(u *url.URL) (Sink, error) {
	_ = "STUB: not implemented"
	return *new(Sink), nil
}

func (sr *sinkRegistry) newFileSinkFromPath(path string) (Sink, error) {
	_ = "STUB: not implemented"
	return *new(Sink), nil
}

func normalizeScheme(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }
