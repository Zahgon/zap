package zapcore

type multiCore []Core

var (
	_ leveledEnabler = multiCore(nil)
	_ Core           = multiCore(nil)
)

func NewTee(cores ...Core) Core { _ = "STUB: not implemented"; return *new(Core) }

func (mc multiCore) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (mc multiCore) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (mc multiCore) Enabled(lvl Level) bool { _ = "STUB: not implemented"; return false }

func (mc multiCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (mc multiCore) Write(ent Entry, fields []Field) error { _ = "STUB: not implemented"; return nil }

func (mc multiCore) Sync() error { _ = "STUB: not implemented"; return nil }
