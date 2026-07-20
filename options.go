package zap

import (
	"go.uber.org/zap/zapcore"
)

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) { _ = "STUB: not implemented"; return }

func WrapCore(f func(zapcore.Core) zapcore.Core) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Hooks(hooks ...func(zapcore.Entry) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Fields(fs ...Field) Option { _ = "STUB: not implemented"; return *new(Option) }

func ErrorOutput(w zapcore.WriteSyncer) Option { _ = "STUB: not implemented"; return *new(Option) }

func Development() Option { _ = "STUB: not implemented"; return *new(Option) }

func AddCaller() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCaller(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func AddCallerSkip(skip int) Option { _ = "STUB: not implemented"; return *new(Option) }

func AddStacktrace(lvl zapcore.LevelEnabler) Option { _ = "STUB: not implemented"; return *new(Option) }

func IncreaseLevel(lvl zapcore.LevelEnabler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPanicHook(hook zapcore.CheckWriteHook) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func OnFatal(action zapcore.CheckWriteAction) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFatalHook(hook zapcore.CheckWriteHook) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithClock(clock zapcore.Clock) Option { _ = "STUB: not implemented"; return *new(Option) }
