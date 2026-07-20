package observer

import "go.uber.org/zap/zapcore"

type LoggedEntry struct {
	zapcore.Entry
	Context []zapcore.Field
}

func (e LoggedEntry) ContextMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }
