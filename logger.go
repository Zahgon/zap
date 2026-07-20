package zap

import (
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	core zapcore.Core

	development bool
	addCaller   bool
	onPanic     zapcore.CheckWriteHook
	onFatal     zapcore.CheckWriteHook

	name        string
	errorOutput zapcore.WriteSyncer

	addStack zapcore.LevelEnabler

	callerSkip int

	clock zapcore.Clock
}

func New(core zapcore.Core, options ...Option) *Logger { _ = "STUB: not implemented"; return nil }

func NewNop() *Logger { _ = "STUB: not implemented"; return nil }

func NewProduction(options ...Option) (*Logger, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDevelopment(options ...Option) (*Logger, error) { _ = "STUB: not implemented"; return nil, nil }

func Must(logger *Logger, err error) *Logger { _ = "STUB: not implemented"; return nil }

func NewExample(options ...Option) *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) Sugar() *SugaredLogger { _ = "STUB: not implemented"; return nil }

func (log *Logger) Named(s string) *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) WithOptions(opts ...Option) *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) With(fields ...Field) *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) WithLazy(fields ...Field) *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) Level() zapcore.Level { _ = "STUB: not implemented"; return *new(zapcore.Level) }

func (log *Logger) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (log *Logger) Log(lvl zapcore.Level, msg string, fields ...Field) {
	_ = "STUB: not implemented"
	return
}

func (log *Logger) Debug(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Info(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Warn(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Error(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) DPanic(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Panic(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Fatal(msg string, fields ...Field) { _ = "STUB: not implemented"; return }

func (log *Logger) Sync() error { _ = "STUB: not implemented"; return nil }

func (log *Logger) Core() zapcore.Core { _ = "STUB: not implemented"; return *new(zapcore.Core) }

func (log *Logger) Name() string { _ = "STUB: not implemented"; return "" }

func (log *Logger) clone() *Logger { _ = "STUB: not implemented"; return nil }

func (log *Logger) check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func terminalHookOverride(defaultHook, override zapcore.CheckWriteHook) zapcore.CheckWriteHook {
	_ = "STUB: not implemented"
	return *new(zapcore.CheckWriteHook)
}
