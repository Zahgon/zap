package zapcore

type Core interface {
	LevelEnabler

	With([]Field) Core

	Check(Entry, *CheckedEntry) *CheckedEntry

	Write(Entry, []Field) error

	Sync() error
}

type nopCore struct{}

func NewNopCore() Core              { _ = "STUB: not implemented"; return *new(Core) }
func (nopCore) Enabled(Level) bool  { _ = "STUB: not implemented"; return false }
func (n nopCore) With([]Field) Core { _ = "STUB: not implemented"; return *new(Core) }
func (nopCore) Check(_ Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}
func (nopCore) Write(Entry, []Field) error { _ = "STUB: not implemented"; return nil }
func (nopCore) Sync() error                { _ = "STUB: not implemented"; return nil }

func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core {
	_ = "STUB: not implemented"
	return *new(Core)
}

type ioCore struct {
	LevelEnabler
	enc Encoder
	out WriteSyncer
}

var (
	_ Core           = (*ioCore)(nil)
	_ leveledEnabler = (*ioCore)(nil)
)

func (c *ioCore) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (c *ioCore) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (c *ioCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *ioCore) Write(ent Entry, fields []Field) error { _ = "STUB: not implemented"; return nil }

func (c *ioCore) Sync() error { _ = "STUB: not implemented"; return nil }

func (c *ioCore) clone() *ioCore { _ = "STUB: not implemented"; return nil }
