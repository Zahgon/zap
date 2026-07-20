package zaptest

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOption interface {
	applyLoggerOption(*loggerOptions)
}

type loggerOptions struct {
	Level      zapcore.LevelEnabler
	zapOptions []zap.Option
}

type loggerOptionFunc func(*loggerOptions)

func (f loggerOptionFunc) applyLoggerOption(opts *loggerOptions) { _ = "STUB: not implemented"; return }

func Level(enab zapcore.LevelEnabler) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

func WrapOptions(zapOpts ...zap.Option) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

func NewLogger(t TestingT, opts ...LoggerOption) *zap.Logger { _ = "STUB: not implemented"; return nil }

type TestingWriter struct {
	t TestingT

	markFailed bool
}

func NewTestingWriter(t TestingT) TestingWriter {
	_ = "STUB: not implemented"
	return *new(TestingWriter)
}

func (w TestingWriter) WithMarkFailed(v bool) TestingWriter {
	_ = "STUB: not implemented"
	return *new(TestingWriter)
}

func (w TestingWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w TestingWriter) Sync() error { _ = "STUB: not implemented"; return nil }
