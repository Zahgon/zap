package zap

import (
	"go.uber.org/zap/zapcore"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	_multipleErrMsg     = "Multiple errors without a key."
)

type SugaredLogger struct {
	base *Logger
}

func (s *SugaredLogger) Desugar() *Logger { _ = "STUB: not implemented"; return nil }

func (s *SugaredLogger) Named(name string) *SugaredLogger { _ = "STUB: not implemented"; return nil }

func (s *SugaredLogger) WithOptions(opts ...Option) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

func (s *SugaredLogger) WithLazy(args ...interface{}) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

func (s *SugaredLogger) Level() zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func (s *SugaredLogger) Log(lvl zapcore.Level, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Debug(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Warn(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) DPanic(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Panic(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Logf(lvl zapcore.Level, template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Debugf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Infof(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Warnf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Errorf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) DPanicf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Panicf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Logln(lvl zapcore.Level, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) Debugln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Infoln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Warnln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Errorln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) DPanicln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Panicln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (s *SugaredLogger) Sync() error { _ = "STUB: not implemented"; return nil }

func (s *SugaredLogger) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *SugaredLogger) logln(lvl zapcore.Level, fmtArgs []interface{}, context []interface{}) {
	_ = "STUB: not implemented"
	return
}

func getMessage(template string, fmtArgs []interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func getMessageln(fmtArgs []interface{}) string { _ = "STUB: not implemented"; return "" }

func (s *SugaredLogger) sweetenFields(args []interface{}) []Field {
	_ = "STUB: not implemented"
	return nil
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
