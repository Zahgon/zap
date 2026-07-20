package zap

import (
	"log"
	"sync"

	"go.uber.org/zap/zapcore"
)

const (
	_stdLogDefaultDepth      = 1
	_loggerWriterDepth       = 2
	_programmerErrorTemplate = "You've found a bug in zap! Please file a bug at " +
		"https://github.com/uber-go/zap/issues/new and reference this error: %v"
)

var (
	_globalMu sync.RWMutex
	_globalL  = NewNop()
	_globalS  = _globalL.Sugar()
)

func L() *Logger { _ = "STUB: not implemented"; return nil }

func S() *SugaredLogger { _ = "STUB: not implemented"; return nil }

func ReplaceGlobals(logger *Logger) func() { _ = "STUB: not implemented"; return nil }

func NewStdLog(l *Logger) *log.Logger { _ = "STUB: not implemented"; return nil }

func NewStdLogAt(l *Logger, level zapcore.Level) (*log.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RedirectStdLog(l *Logger) func() { _ = "STUB: not implemented"; return nil }

func RedirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func redirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func levelToFunc(logger *Logger, lvl zapcore.Level) (func(string, ...Field), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type loggerWriter struct {
	logFunc func(msg string, fields ...Field)
}

func (l *loggerWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
