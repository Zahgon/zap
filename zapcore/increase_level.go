package zapcore

type levelFilterCore struct {
	core  Core
	level LevelEnabler
}

var (
	_ Core           = (*levelFilterCore)(nil)
	_ leveledEnabler = (*levelFilterCore)(nil)
)

func NewIncreaseLevelCore(core Core, level LevelEnabler) (Core, error) {
	_ = "STUB: not implemented"
	return *new(Core), nil
}

func (c *levelFilterCore) Enabled(lvl Level) bool { _ = "STUB: not implemented"; return false }

func (c *levelFilterCore) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (c *levelFilterCore) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (c *levelFilterCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *levelFilterCore) Write(ent Entry, fields []Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *levelFilterCore) Sync() error { _ = "STUB: not implemented"; return nil }
