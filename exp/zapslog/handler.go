//go:build go1.21

package zapslog

import (
	"context"
	"log/slog"

	"go.uber.org/zap/zapcore"
)

type Handler struct {
	core       zapcore.Core
	name       string
	addCaller  bool
	addStackAt slog.Level
	callerSkip int

	groups []string
}

func NewHandler(core zapcore.Core, opts ...HandlerOption) *Handler {
	_ = "STUB: not implemented"
	return nil
}

var _ slog.Handler = (*Handler)(nil)

type groupObject []slog.Attr

func (gs groupObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func convertAttrToField(attr slog.Attr) zapcore.Field {
	_ = "STUB: not implemented"
	return *new(zapcore.Field)
}

func convertSlogLevel(l slog.Level) zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *Handler) Handle(_ context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) appendGroups(fields []zapcore.Field) []zapcore.Field {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *Handler) WithGroup(group string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}
