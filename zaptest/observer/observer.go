package observer

import (
	"sync"

	"go.uber.org/zap/internal"
	"go.uber.org/zap/zapcore"
)

type ObservedLogs struct {
	mu   sync.RWMutex
	logs []LoggedEntry
}

func (o *ObservedLogs) Len() int { _ = "STUB: not implemented"; return 0 }

func (o *ObservedLogs) All() []LoggedEntry { _ = "STUB: not implemented"; return nil }

func (o *ObservedLogs) TakeAll() []LoggedEntry { _ = "STUB: not implemented"; return nil }

func (o *ObservedLogs) AllUntimed() []LoggedEntry { _ = "STUB: not implemented"; return nil }

func (o *ObservedLogs) FilterLevelExact(level zapcore.Level) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) FilterMessage(msg string) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) FilterLoggerName(name string) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) FilterMessageSnippet(snippet string) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) FilterField(field zapcore.Field) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) FilterFieldKey(key string) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) Filter(keep func(LoggedEntry) bool) *ObservedLogs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObservedLogs) add(log LoggedEntry) { _ = "STUB: not implemented"; return }

func New(enab zapcore.LevelEnabler) (zapcore.Core, *ObservedLogs) {
	_ = "STUB: not implemented"
	return *new(zapcore.Core), nil
}

type contextObserver struct {
	zapcore.LevelEnabler
	logs    *ObservedLogs
	context []zapcore.Field
}

var (
	_ zapcore.Core            = (*contextObserver)(nil)
	_ internal.LeveledEnabler = (*contextObserver)(nil)
)

func (co *contextObserver) Level() zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func (co *contextObserver) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (co *contextObserver) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (co *contextObserver) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (co *contextObserver) Sync() error { _ = "STUB: not implemented"; return nil }
