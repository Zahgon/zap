//go:build go1.21

package zapslog

import "log/slog"

type HandlerOption interface {
	apply(*Handler)
}

type handlerOptionFunc func(*Handler)

func (f handlerOptionFunc) apply(handler *Handler) { _ = "STUB: not implemented"; return }

func WithName(name string) HandlerOption { _ = "STUB: not implemented"; return *new(HandlerOption) }

func WithCaller(enabled bool) HandlerOption { _ = "STUB: not implemented"; return *new(HandlerOption) }

func WithCallerSkip(skip int) HandlerOption { _ = "STUB: not implemented"; return *new(HandlerOption) }

func AddStacktraceAt(lvl slog.Level) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}
